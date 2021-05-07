package matchesmanager

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/centrifugal/centrifuge"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/token"
)

var (
	ErrMatchNotFound = errors.New("match not found")
	ErrAllKeysGiven  = errors.New("all player keys were given")
)

type Manager struct {
	node    *centrifuge.Node
	storage *gamesstorage.Storage

	mu      sync.RWMutex
	matches map[string]*match
}

type MatchInfo struct {
	GameID int32
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
		matches: make(map[string]*match),
	}, nil
}

func (m *Manager) NewPlayerKey(session string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if ma, ok := m.matches[session]; ok {
		key, keyOk := ma.newPlayerKey()
		if !keyOk {
			return "", ErrAllKeysGiven
		}
		return key, nil
	}
	return "", ErrMatchNotFound
}

func (m *Manager) MatchInfo(session string) (*MatchInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if ma, ok := m.matches[session]; ok {
		return &MatchInfo{
			GameID: ma.gameID,
		}, nil
	}
	return nil, ErrMatchNotFound
}

func (m *Manager) StartMatch(ctx context.Context, gameID int32) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	ma, err := newMatch(ctx, m.storage, gameID)
	if err != nil {
		return "", fmt.Errorf("can't create new match: %w", err)
	}
	tk, err := m.generateMatchToken(ctx, ma)
	if err != nil {
		return "", fmt.Errorf("can't generate match token: %w", err)
	}
	return tk, nil
}

func (m *Manager) generateMatchToken(ctx context.Context, ma *match) (string, error) {
	for {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		tk := token.New()
		if _, exists := m.matches[tk]; !exists {
			m.matches[tk] = ma
			return tk, nil
		}
	}
}
