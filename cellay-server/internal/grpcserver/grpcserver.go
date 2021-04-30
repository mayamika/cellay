package grpcserver

import (
	"context"
	"fmt"
	"net"
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const shutdownTimeout = 5 * time.Second

type Config struct {
	Addr string
}

type Params struct {
	fx.In

	Logger     *zap.Logger
	LC         fx.Lifecycle
	Shutdowner fx.Shutdowner
	Config     Config
}

func New(p Params) *grpc.Server {
	logger := p.Logger.Named("grpc")
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpczap.UnaryServerInterceptor(logger.Named("unary")),
			grpcrecovery.UnaryServerInterceptor(newPanicHandler(logger)),
		)),
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpczap.StreamServerInterceptor(logger.Named("stream")),
			grpcrecovery.StreamServerInterceptor(newPanicHandler(logger)),
		)),
	)
	reflection.Register(server)
	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", p.Config.Addr)
			if err != nil {
				return fmt.Errorf("can't listen tcp: %w", err)
			}
			go func() {
				defer func() { _ = p.Shutdowner.Shutdown() }()
				if err := server.Serve(listener); err != nil {
					logger.Error("serve grpc failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
			defer cancel()
			closed := make(chan struct{})
			defer close(closed)
			go func() {
				select {
				case <-closed:
					return
				case <-ctx.Done():
					server.Stop()
				}
			}()
			server.GracefulStop()
			return nil
		},
	})
	return server
}

func newPanicHandler(logger *zap.Logger) grpcrecovery.Option {
	return grpcrecovery.WithRecoveryHandler(
		func(reason interface{}) error {
			logger.Error(
				"panic",
				zap.Reflect("reason", reason),
				zap.Stack("stack"),
			)
			return nil
		},
	)
}
