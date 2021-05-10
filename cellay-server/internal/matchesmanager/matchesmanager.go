package matchesmanager

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/centrifugal/centrifuge"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager/game"
	"github.com/mayamika/cellay/cellay-server/internal/token"
)

var (
	ErrMatchNotFound = errors.New("match not found")
	ErrAllKeysGiven  = errors.New("all player keys were given")
)

type Manager struct {
	node      *centrifuge.Node
	wsHandler http.Handler
	storage   *gamesstorage.Storage

	mu         sync.RWMutex
	matches    map[string]*match
	playerKeys map[string]string
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
	m := &Manager{
		node:      node,
		wsHandler: centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{}),
		storage:   p.Storage,
		matches:   make(map[string]*match),
	}
	node.OnConnecting(m.onConnecting)
	node.OnConnect(m.onConnect)
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
	return m, nil
}

func (m *Manager) WebsocketHandler() http.Handler {
	return m.wsHandler
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
	session, err := m.newMatch(ctx, gameID)
	if err != nil {
		return "", fmt.Errorf("can't create new match: %w", err)
	}
	return session, nil
}

func (m *Manager) onConnecting(
	ctx context.Context,
	event centrifuge.ConnectEvent,
) (centrifuge.ConnectReply, error) {
	return centrifuge.ConnectReply{
		Credentials: &centrifuge.Credentials{
			UserID: "awd",
		},
	}, nil
}

func (m *Manager) onConnect(client *centrifuge.Client) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	session := m.playerKeys[client.UserID()]
	player := m.matches[session].checkPlayerKey(client.UserID())
	if player == 0 {
		client.Disconnect(centrifuge.DisconnectBadRequest)
		return
	}
	client.OnMessage(func(me centrifuge.MessageEvent) {
	})
	if err := client.Subscribe(session); err != nil {
		client.Disconnect(centrifuge.DisconnectBadRequest)
	}
}

func (m *Manager) generateMatchSession(ctx context.Context) (string, error) {
	for {
		if err := ctx.Err(); err != nil {
			return "", err
		}
		session := token.New()
		if _, exists := m.matches[session]; !exists {
			return session, nil
		}
	}
}

func (m *Manager) generatePlayerKeys(ctx context.Context, players int) ([]string, error) {
	newKeys := make(map[string]struct{})
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		key := token.New()
		if _, exists := m.playerKeys[key]; !exists {
			newKeys[key] = struct{}{}
		}
		if len(newKeys) == players {
			break
		}
	}
	var keys []string
	for key := range newKeys {
		keys = append(keys, key)
	}
	return keys, nil
}

func newGame(ctx context.Context, storage *gamesstorage.Storage, gameID int32) (*game.Game, error) {
	code, err := storage.GameCode(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game code: %w", err)
	}
	assets, err := storage.GameAssets(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game assets: %w", err)
	}
	layers := make([]string, 0)
	for name := range assets.Layers {
		layers = append(layers, name)
	}
	return game.New(&game.Config{
		Code: code.Code,
		Field: game.Field{
			Cols: int(assets.Field.Cols),
			Rows: int(assets.Field.Rows),
		},
		Layers: layers,
	})
}

func (m *Manager) newMatch(ctx context.Context, gameID int32) (string, error) {
	g, err := newGame(ctx, m.storage, gameID)
	if err != nil {
		return "", fmt.Errorf("can't create game: %w", err)
	}
	session, err := m.generateMatchSession(ctx)
	if err != nil {
		return "", fmt.Errorf("can't generate match session: %w", err)
	}
	keys, err := m.generatePlayerKeys(ctx, 2)
	if err != nil {
		return "", fmt.Errorf("can't generate match player keys: %w", err)
	}
	m.matches[session] = &match{
		gameID: gameID,
		game:   g,
		keys:   keys,
	}
	for _, key := range keys {
		m.playerKeys[key] = session
	}
	return session, nil
}
