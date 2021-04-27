package gamesstorage

import (
	"context"
	"fmt"

	"github.com/genjidb/genji"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Storage struct {
	db *genji.DB
}

type Config struct {
	Path string
}

type Params struct {
	fx.In

	Logger *zap.Logger
	LC     fx.Lifecycle
	Config Config
}

func New(p Params) *Storage {
	s := &Storage{}
	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			s.db, err = genji.Open(p.Config.Path)
			if err != nil {
				return fmt.Errorf("can't open genji db: %w", err)
			}
			if err := applyMigrations(ctx, s.db); err != nil {
				return fmt.Errorf("can't apply db migrations: %w", err)
			}
			return nil
		},
		OnStop: func(_ context.Context) error {
			return s.db.Close()
		},
	})
	return s
}

type GameInfo struct {
	ID          int32
	Name        string
	Description string
}

type GameCode struct {
	ID   int32
	Code string
}

type GameAssetsField struct {
	Rows int32
	Cols int32
}

type GameAssetsLayer struct {
	Width   int32
	Height  int32
	Depth   int32
	Texture []byte
}

type GameAssets struct {
	ID     int32
	Field  GameAssetsField
	Layers map[string]*GameAssetsLayer
}

type game struct {
	ID          int32                       `genji:"id"`
	Name        string                      `genji:"name"`
	Description string                      `genji:"description"`
	Code        string                      `genji:"code"`
	Field       GameAssetsField             `genji:"field"`
	Layers      map[string]*GameAssetsLayer `genji:"layers"`
}

func (s *Storage) GameInfo(ctx context.Context, id int32) (*GameInfo, error) {
	return nil, nil
}

func (s *Storage) GameCode(ctx context.Context, id int32) (*GameInfo, error) {
	return nil, nil
}

func (s *Storage) GameAssets(ctx context.Context, id int32) (*GameInfo, error) {
	return nil, nil
}
