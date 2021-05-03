package matchesmanager

import (
	"context"
	"fmt"
	"sync"

	"github.com/centrifugal/centrifuge"
	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/token"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Manager struct {
	node    *centrifuge.Node
	storage *gamesstorage.Storage

	mu          sync.Mutex
	matchTokens map[string]struct{} //nolint:unused // Not implemented
}

type Config struct {
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
	_, err := m.storage.GameCode(ctx, gameID)
	if err != nil {
		return "", fmt.Errorf("can't fetch game code: %w", err)
	}
	return "", nil
}

//nolint:unused // Not implemented
func (m *Manager) newMatchToken(ctx context.Context) (string, error) {
	for {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		t := token.New()
		if _, exists := m.matchTokens[t]; !exists {
			m.matchTokens[t] = struct{}{}
			return t, nil
		}
	}
}
