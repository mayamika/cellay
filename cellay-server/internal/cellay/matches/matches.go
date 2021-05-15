package matches

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mayamika/cellay/cellay-server/internal/httpserver"
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager"
	cellayv1 "github.com/mayamika/cellay/proto/cellay/v1"
)

type Service struct {
	cellayv1.UnimplementedMatchesServiceServer
	manager *matchesmanager.Manager
}

type Params struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *httpserver.Server
	Manager    *matchesmanager.Manager
}

func NewService(p Params) *Service {
	s := &Service{
		manager: p.Manager,
	}
	cellayv1.RegisterMatchesServiceServer(p.GRPCServer, s)
	p.HTTPServer.RegisterService(cellayv1.RegisterMatchesServiceHandler)
	return s
}

func (s *Service) New(
	ctx context.Context,
	req *cellayv1.MatchesServiceNewRequest,
) (*cellayv1.MatchesServiceNewResponse, error) {
	session, err := s.manager.StartMatch(ctx, req.GetGameId())
	if err != nil {
		return nil, errInternalf("can't start new match: %v", err)
	}
	key, err := s.manager.NewPlayerKey(session)
	if err != nil {
		return nil, errInternalf("can't get player key: %v", err)
	}
	return &cellayv1.MatchesServiceNewResponse{
		Session: session,
		Key:     key,
	}, nil
}

func (s *Service) Info(
	ctx context.Context,
	req *cellayv1.MatchesServiceInfoRequest,
) (*cellayv1.MatchesServiceInfoResponse, error) {
	info, err := s.manager.MatchInfo(req.GetSession())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "can't get match info: %v", err)
	}
	res := &cellayv1.MatchesServiceInfoResponse{
		GameId: info.GameID,
	}
	if req.GetNew() {
		res.Key, err = s.manager.NewPlayerKey(req.GetSession())
		if err != nil {
			return nil, errInternalf("can't get player key: %v", err)
		}
	}
	return res, nil
}

func errInternalf(format string, args ...interface{}) error {
	return status.Errorf(codes.Internal, format, args...)
}
