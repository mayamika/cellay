package gamesstorage

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

func TestStorage(t *testing.T) {
	t.Run("InsertGet", withStorage(testInsertGet))
}

func testInsertGet(t *testing.T, s *Storage) {
	var (
		r   = require.New(t)
		ctx = newTestCtx(t)
	)
	firstGame := newTestGame(1)
	r.NoError(s.AddGame(ctx, firstGame))
	secondGame := newTestGame(2)
	r.NoError(s.AddGame(ctx, secondGame))
	games, err := s.AllGames(ctx)
	r.NoError(err)
	r.Len(games, 2)
	r.Equal(&GameInfo{
		ID:          1,
		Name:        firstGame.Name,
		Description: firstGame.Description,
	}, games[0])
	r.Equal(&GameInfo{
		ID:          2,
		Name:        secondGame.Name,
		Description: secondGame.Description,
	}, games[1])

	info, err := s.GameInfo(ctx, 1)
	r.NoError(err)
	r.Equal(&GameInfo{
		ID:          1,
		Name:        firstGame.Name,
		Description: firstGame.Description,
	}, info)

	code, err := s.GameCode(ctx, 1)
	r.NoError(err)
	r.Equal(&GameCode{
		ID:   1,
		Code: firstGame.Code,
	}, code)

	assets, err := s.GameAssets(ctx, 1)
	r.NoError(err)
	r.Equal(&GameAssets{
		ID:     1,
		Field:  firstGame.Field,
		Layers: firstGame.Layers,
	}, assets)
}

func withStorage(fn func(t *testing.T, s *Storage)) func(t *testing.T) {
	return func(t *testing.T) {
		var s *Storage
		app := fxtest.New(t,
			fx.Provide(New),
			fx.Supply(newTestLogger(t)),
			fx.Supply(Config{
				Path: filepath.Join(t.TempDir(), "db"),
			}),
			fx.Populate(&s),
		)
		app.RequireStart()
		fn(t, s)
		app.RequireStop()
	}
}

func newTestLogger(t *testing.T) *zap.Logger {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	return logger
}

func newTestCtx(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

func newTestGame(equivalence int) *Game {
	return &Game{
		Name:        fmt.Sprintf("name.%d", equivalence),
		Description: fmt.Sprintf("desc.%d", equivalence),
		Code:        fmt.Sprintf("code.%d", equivalence),
		Field: GameAssetsField{
			Rows: int32(equivalence),
			Cols: int32(equivalence),
		},
		Layers: map[string]*GameAssetsLayer{
			fmt.Sprint(equivalence): {
				Width:   int32(equivalence),
				Height:  int32(equivalence),
				Depth:   int32(equivalence),
				Texture: []byte(fmt.Sprint(equivalence)),
			},
		},
	}
}
