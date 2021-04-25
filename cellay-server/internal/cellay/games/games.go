package games

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
	cellayv1.UnimplementedGamesServiceServer
}

type Params struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *httpserver.Server
}

func NewService(p Params) *Service {
	s := &Service{}
	cellayv1.RegisterGamesServiceServer(p.GRPCServer, s)
	p.HTTPServer.RegisterService(cellayv1.RegisterGamesServiceHandler)
	return s
}

func (s *Service) GetInfo(
	ctx context.Context,
	req *cellayv1.GamesServiceGetInfoRequest,
) (*cellayv1.GamesServiceGetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func (s *Service) GetCode(
	ctx context.Context,
	req *cellayv1.GamesServiceGetCodeRequest,
) (*cellayv1.GamesServiceGetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}

func (s *Service) GetAssets(
	ctx context.Context,
	req *cellayv1.GamesServiceGetAssetsRequest,
) (*cellayv1.GamesServiceGetAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}

func (s *Service) GetAll(
	ctx context.Context,
	req *cellayv1.GamesServiceGetAllRequest,
) (*cellayv1.GamesServiceGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
