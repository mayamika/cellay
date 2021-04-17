package app

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/mayamika/cellay/cellay-server/internal/cellay/games"
	"github.com/mayamika/cellay/cellay-server/internal/cellay/matches"
	"github.com/mayamika/cellay/cellay-server/internal/grpcserver"
	"github.com/mayamika/cellay/cellay-server/internal/httpserver"
	"github.com/mayamika/cellay/cellay-server/internal/logger"
)

type Config struct {
	fx.Out

	GRPC grpcserver.Config
	HTTP httpserver.Config
}

func NewDefaultConfig() *Config {
	return &Config{
		GRPC: grpcserver.Config{
			Addr: ":8081",
		},
		HTTP: httpserver.Config{
			Addr: ":8080",
		},
	}
}

func New(config *Config) *fx.App {
	if config == nil {
		config = NewDefaultConfig()
	}
	return fx.New(
		fx.Supply(*config),
		fx.Provide(logger.New),
		fx.Provide(grpcserver.New),
		fx.Provide(httpserver.New),
		fx.Provide(games.NewService),
		fx.Provide(matches.NewService),
		fx.Invoke(onStart),
	)
}

func onStart(
	_ *grpc.Server,
	_ *httpserver.Server,
	_ *games.Service,
	_ *matches.Service,
) {
}
