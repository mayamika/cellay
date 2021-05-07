package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/mayamika/cellay/cellay-server/internal/grpcserver"
)

type ServiceRegisterFunc func(ctx context.Context, gwMux *runtime.ServeMux, conn *grpc.ClientConn) error

type Server struct {
	mux                  *http.ServeMux
	grpcClient           *grpc.ClientConn
	serviceRegisterFuncs []ServiceRegisterFunc
}

type Config struct {
	Addr string
}

type Params struct {
	fx.In

	Logger     *zap.Logger
	LC         fx.Lifecycle
	Shutdowner fx.Shutdowner
	Config     Config
	GRPCConfig grpcserver.Config
}

func New(p Params) *Server { //nolint: gocritic
	var (
		logger = p.Logger.Named("http")
		server = &Server{
			mux: http.NewServeMux(),
		}
		httpServer = &http.Server{
			Addr:    p.Config.Addr,
			Handler: server.mux,
		}
	)
	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			server.grpcClient, err = grpc.DialContext(
				ctx,
				p.GRPCConfig.Addr,
				grpc.WithInsecure(),
			)
			if err != nil {
				return fmt.Errorf("can't dial grpc server: %w", err)
			}
			gwMux := runtime.NewServeMux()
			for _, register := range server.serviceRegisterFuncs {
				if err := register(ctx, gwMux, server.grpcClient); err != nil {
					return fmt.Errorf("can't register service: %w", err)
				}
			}
			server.mux.Handle(`/api/v1/`, http.StripPrefix(`/api/v1`, gwMux))
			go func() {
				defer func() { _ = p.Shutdowner.Shutdown() }()
				if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					logger.Error("serve http failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			defer func() { _ = server.grpcClient.Close() }()
			if err := httpServer.Shutdown(ctx); err != nil {
				return fmt.Errorf("http server shutdown failed: %w", err)
			}
			return nil
		},
	})
	return server
}

func (s *Server) Handle(path string, handler http.Handler) {
	s.mux.Handle(path, handler)
}

func (s *Server) RegisterService(funcs ...ServiceRegisterFunc) {
	s.serviceRegisterFuncs = append(s.serviceRegisterFuncs, funcs...)
}
