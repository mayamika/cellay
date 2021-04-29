package gamesstorage

import (
	"context"
	"fmt"

	"github.com/genjidb/genji"
	"github.com/genjidb/genji/document"
	"go.uber.org/fx"
	"go.uber.org/multierr"
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
	Field  GameAssetsField `genji:"game_field"`
	Layers map[string]*GameAssetsLayer
}

type Game struct {
	ID          int32
	Name        string
	Description string
	Code        string
	Field       GameAssetsField `genji:"game_field"`
	Layers      map[string]*GameAssetsLayer
}

func (s *Storage) GameInfo(ctx context.Context, id int32) (*GameInfo, error) {
	db := s.db.WithContext(ctx)
	res, err := db.QueryDocument(`SELECT pk(), name, description FROM games WHERE pk() = ?`, id)
	if err != nil {
		return nil, fmt.Errorf("query document failed: %w", err)
	}
	info := &GameInfo{}
	if err := scanStructAndID(res, info, &info.ID); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Storage) GameCode(ctx context.Context, id int32) (*GameCode, error) {
	db := s.db.WithContext(ctx)
	res, err := db.QueryDocument(`SELECT pk(), code FROM games WHERE pk() = ?`, id)
	if err != nil {
		return nil, fmt.Errorf("query document failed: %w", err)
	}
	code := &GameCode{}
	if err := scanStructAndID(res, code, &code.ID); err != nil {
		return nil, err
	}
	return code, nil
}

func (s *Storage) GameAssets(ctx context.Context, id int32) (*GameAssets, error) {
	db := s.db.WithContext(ctx)
	res, err := db.QueryDocument(`SELECT pk(), game_field, layers FROM games WHERE pk() = ?`, id)
	if err != nil {
		return nil, fmt.Errorf("query document failed: %w", err)
	}
	assets := &GameAssets{}
	if err := scanStructAndID(res, assets, &assets.ID); err != nil {
		return nil, err
	}
	return assets, nil
}

func (s *Storage) AllGames(ctx context.Context) ([]*GameInfo, error) {
	db := s.db.WithContext(ctx)
	res, err := db.Query(`SELECT pk(), name, description FROM games`)
	if err != nil {
		return nil, fmt.Errorf("can't get query result: %w", err)
	}
	var gameInfos []*GameInfo
	err = res.Iterate(func(d document.Document) error {
		info := &GameInfo{}
		if err := scanStructAndID(d, info, &info.ID); err != nil {
			return err
		}
		gameInfos = append(gameInfos, info)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("can't iterate through query result: %w", err)
	}
	if err := res.Close(); err != nil {
		return nil, fmt.Errorf("can't close query result: %w", err)
	}
	return gameInfos, nil
}

func (s *Storage) AddGame(ctx context.Context, game *Game) error {
	db := s.db.WithContext(ctx)
	if err := db.Exec(`INSERT INTO games VALUES ?`, game); err != nil {
		return fmt.Errorf("can't insert document: %w", err)
	}
	return nil
}

func inTx(ctx context.Context, db *genji.DB, fn func(tx *genji.Tx) error) error {
	db = db.WithContext(ctx)
	tx, err := db.Begin(true)
	if err != nil {
		return fmt.Errorf("can't begin tx: %w", err)
	}
	tryRollback := func(err error) error {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return multierr.Combine(
				fmt.Errorf("can't rollback tx: %w", rollbackErr),
				err,
			)
		}
		return err
	}
	if err := fn(tx); err != nil {
		return tryRollback(fmt.Errorf("can't run tx: %w", err))
	}
	if err := tx.Commit(); err != nil {
		return tryRollback(fmt.Errorf("can't commit tx: %w", err))
	}
	return nil
}

func scanStructAndID(doc document.Document, structPtr interface{}, id *int32) error {
	if err := document.StructScan(doc, structPtr); err != nil {
		return fmt.Errorf("can't scan document: %w", err)
	}
	documentID, err := doc.GetByField("pk()")
	if err != nil {
		return fmt.Errorf("can't get document id: %w", err)
	}
	if err := documentID.Scan(id); err != nil {
		return fmt.Errorf("can't get document id: %w", err)
	}
	return nil
}
