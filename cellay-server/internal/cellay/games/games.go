package games

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/httpserver"
	cellayv1 "github.com/mayamika/cellay/proto/cellay/v1"
)

type Service struct {
	cellayv1.UnimplementedGamesServiceServer
	storage *gamesstorage.Storage
}

type Params struct {
	fx.In

	GRPCServer *grpc.Server
	HTTPServer *httpserver.Server
	Storage    *gamesstorage.Storage
}

func NewService(p Params) *Service {
	s := &Service{
		storage: p.Storage,
	}
	cellayv1.RegisterGamesServiceServer(p.GRPCServer, s)
	p.HTTPServer.RegisterService(cellayv1.RegisterGamesServiceHandler)
	return s
}

func (s *Service) GetInfo(
	ctx context.Context,
	req *cellayv1.GamesServiceGetInfoRequest,
) (*cellayv1.GamesServiceGetInfoResponse, error) {
	info, err := s.storage.GameInfo(ctx, req.GetId())
	if err != nil {
		return nil, errInternalf("can't get game info from storage: %v", err)
	}
	return gameInfoToProto(info), nil
}

func (s *Service) GetCode(
	ctx context.Context,
	req *cellayv1.GamesServiceGetCodeRequest,
) (*cellayv1.GamesServiceGetCodeResponse, error) {
	code, err := s.storage.GameCode(ctx, req.GetId())
	if err != nil {
		return nil, errInternalf("can't get game code from storage: %v", err)
	}
	return &cellayv1.GamesServiceGetCodeResponse{
		Id:   code.ID,
		Code: code.Code,
	}, nil
}

func (s *Service) GetAssets(
	ctx context.Context,
	req *cellayv1.GamesServiceGetAssetsRequest,
) (*cellayv1.GamesServiceGetAssetsResponse, error) {
	assets, err := s.storage.GameAssets(ctx, req.GetId())
	if err != nil {
		return nil, errInternalf("can't get game assets from storage: %v", err)
	}
	return gameAssetsToProto(assets), nil
}

func (s *Service) GetAll(
	ctx context.Context,
	req *cellayv1.GamesServiceGetAllRequest,
) (*cellayv1.GamesServiceGetAllResponse, error) {
	games, err := s.storage.AllGames(ctx)
	if err != nil {
		return nil, errInternalf("can't get games from storage: %v", err)
	}
	var gamesProto []*cellayv1.GamesServiceGetInfoResponse
	for _, info := range games {
		gamesProto = append(gamesProto, gameInfoToProto(info))
	}
	return &cellayv1.GamesServiceGetAllResponse{
		Games: gamesProto,
		Total: int32(len(gamesProto)),
	}, nil
}

func (s *Service) Add(
	ctx context.Context,
	req *cellayv1.GamesServiceAddRequest,
) (*cellayv1.GamesServiceAddResponse, error) {
	if err := s.storage.AddGame(ctx, gameFromProto(req)); err != nil {
		return nil, errInternalf("can't add game to storage: %v", err)
	}
	return &cellayv1.GamesServiceAddResponse{}, nil
}

func errInternalf(format string, args ...interface{}) error {
	return status.Errorf(codes.Internal, format, args...)
}

func gameInfoToProto(info *gamesstorage.GameInfo) *cellayv1.GamesServiceGetInfoResponse {
	return &cellayv1.GamesServiceGetInfoResponse{
		Id:          info.ID,
		Name:        info.Name,
		Description: info.Description,
	}
}

func gameAssetsToProto(assets *gamesstorage.GameAssets) *cellayv1.GamesServiceGetAssetsResponse {
	layers := make(map[string]*cellayv1.GameAssetsLayer)
	for name, layer := range assets.Layers {
		layers[name] = &cellayv1.GameAssetsLayer{
			Width:   layer.Width,
			Height:  layer.Height,
			Depth:   layer.Depth,
			Texture: layer.Texture,
		}
	}
	return &cellayv1.GamesServiceGetAssetsResponse{
		Id: assets.ID,
		Field: &cellayv1.GameAssetsField{
			Rows: assets.Field.Rows,
			Cols: assets.Field.Cols,
		},
		Layers: layers,
		Background: &cellayv1.GameAssetsTexture{
			Width:   assets.Background.Width,
			Height:  assets.Background.Height,
			Texture: assets.Background.Texture,
		},
	}
}

func gameFromProto(req *cellayv1.GamesServiceAddRequest) *gamesstorage.Game {
	layers := make(map[string]*gamesstorage.GameAssetsLayer)
	for name, layer := range req.GetLayers() {
		layers[name] = &gamesstorage.GameAssetsLayer{
			Width:   layer.GetWidth(),
			Height:  layer.GetHeight(),
			Depth:   layer.GetDepth(),
			Texture: layer.GetTexture(),
		}
	}
	return &gamesstorage.Game{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Code:        req.GetCode(),
		Field: gamesstorage.GameAssetsField{
			Rows: req.GetField().GetRows(),
			Cols: req.GetField().GetCols(),
		},
		Background: gamesstorage.GameAssetsTexture{
			Width:   req.GetBackground().GetWidth(),
			Height:  req.GetBackground().GetHeight(),
			Texture: req.GetBackground().GetTexture(),
		},
		Layers: layers,
	}
}
