package matches

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/mayamika/cellay/cellay-server/internal/httpserver"
	cellayv1 "github.com/mayamika/cellay/proto/cellay/v1"
)

type Service struct {
	cellayv1.UnimplementedMatchesServiceServer
}

type Params struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *httpserver.Server
}

func NewService(p Params) *Service {
	s := &Service{}
	cellayv1.RegisterMatchesServiceServer(p.GRPCServer, s)
	p.HTTPServer.RegisterService(cellayv1.RegisterMatchesServiceHandler)
	return s
}
