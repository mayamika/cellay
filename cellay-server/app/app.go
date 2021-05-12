package app

import (
	"fmt"
	"os"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/mayamika/cellay/cellay-server/internal/cellay/games"
	"github.com/mayamika/cellay/cellay-server/internal/cellay/matches"
	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/grpcserver"
	"github.com/mayamika/cellay/cellay-server/internal/httpserver"
	"github.com/mayamika/cellay/cellay-server/internal/logger"
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager"
)

type Config struct {
	fx.Out

	GRPC    grpcserver.Config
	HTTP    httpserver.Config
	Storage gamesstorage.Config
}

func NewDefaultConfig() *Config {
	return &Config{
		GRPC: grpcserver.Config{
			Addr: ":8081",
		},
		HTTP: httpserver.Config{
			Addr: ":8080",
		},
		Storage: gamesstorage.Config{
			Path: "gamesstorage.db",
		},
	}
}

func ParseFlagsAndConfig() (*Config, error) {
	config := NewDefaultConfig()
	fmt.Fprint(os.Stdout, spew.Sdump(config))
	return config, nil
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
		fx.Provide(gamesstorage.New),
		fx.Provide(matchesmanager.New),
		fx.Invoke(onStart),
	)
}

func onStart(
	_ *grpc.Server,
	httpServer *httpserver.Server,
	_ *games.Service,
	_ *matches.Service,
	_ *gamesstorage.Storage,
	matchesManager *matchesmanager.Manager,
) {
	httpServer.Handle(`/connection/websocket`, matchesManager.WebsocketHandler())
	fmt.Printf("%+v", matchesManager.WebsocketHandler())
}
