package matchesmanager

import (
	"context"
	"fmt"
	"sync"

	"github.com/centrifugal/centrifuge"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager/match"
	"github.com/mayamika/cellay/cellay-server/internal/token"
)

type Manager struct {
	node    *centrifuge.Node
	storage *gamesstorage.Storage

	mu      sync.Mutex
	matches map[string]*match.Match //nolint:unused // Not implemented
}

type Params struct {
	fx.In

	Logger  *zap.Logger
	LC      fx.Lifecycle
	Storage *gamesstorage.Storage
}

func New(p Params) (*Manager, error) {
	node, err := centrifuge.New(centrifuge.DefaultConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create centrifuge node: %w", err)
	}
	p.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			if err := node.Run(); err != nil {
				return fmt.Errorf("can't start centrifuge node: %w", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := node.Shutdown(ctx); err != nil {
				return fmt.Errorf("can't shutdown centrifuge node: %w", err)
			}
			return nil
		},
	})
	return &Manager{
		node:    node,
		storage: p.Storage,
	}, nil
}

func (m *Manager) StartMatch(ctx context.Context, gameID int32) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	match, err := m.newMatch(ctx, gameID)
	if err != nil {
		return "", fmt.Errorf("can't create new match: %w", err)
	}
	token, err := m.generateMatchToken(ctx, match)
	if err != nil {
		return "", fmt.Errorf("can't generate match token: %w", err)
	}
	return token, nil
}

func (m *Manager) newMatch(ctx context.Context, gameID int32) (*match.Match, error) {
	code, err := m.storage.GameCode(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game code: %w", err)
	}
	assets, err := m.storage.GameAssets(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game assets: %w", err)
	}
	layers := make([]string, 0)
	for name := range assets.Layers {
		layers = append(layers, name)
	}
	return match.New(&match.Config{
		Code: code.Code,
		Field: match.Field{
			Cols: int(assets.Field.Cols),
			Rows: int(assets.Field.Rows),
		},
		Layers: layers,
	})
}

func (m *Manager) generateMatchToken(ctx context.Context, match *match.Match) (string, error) {
	for {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		t := token.New()
		if _, exists := m.matches[t]; !exists {
			m.matches[t] = match
			return t, nil
		}
	}
}
