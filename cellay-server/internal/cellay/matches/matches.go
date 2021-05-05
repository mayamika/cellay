package matches

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (s *Service) New(
	ctx context.Context,
	req *cellayv1.MatchesServiceNewRequest,
) (*cellayv1.MatchesServiceNewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}

func (s *Service) Info(
	ctx context.Context,
	req *cellayv1.MatchesServiceInfoRequest,
) (*cellayv1.MatchesServiceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
