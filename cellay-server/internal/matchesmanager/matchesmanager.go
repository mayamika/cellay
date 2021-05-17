package matchesmanager

import (
	"context"
	"encoding/json"
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
	ErrInvalidPlayerKey = errors.New("invalid player key")
	ErrMatchNotFound    = errors.New("match not found")
	ErrAllKeysGiven     = errors.New("all player keys were given")
)

type Manager struct {
	logger    *zap.Logger
	node      *centrifuge.Node
	wsHandler http.Handler
	storage   *gamesstorage.Storage

	mu             sync.RWMutex
	matches        map[string]*match
	playerSessions map[string]string
}

type PlayerInfo struct {
	PlayerID int32
}

type MatchInfo struct {
	GameID   int32
	GameName string
}

const (
	actionTypeClick = "click"
	actionTypeMove  = "move"
)

type actionMessage struct {
	Type     string
	X, Y     int
	From, To struct {
		X, Y int
	}
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
		logger:         p.Logger.Named("matchesmanager"),
		node:           node,
		wsHandler:      centrifuge.NewWebsocketHandler(node, newWebsocketConfig()),
		storage:        p.Storage,
		matches:        make(map[string]*match),
		playerSessions: make(map[string]string),
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
	ma, ok := m.matches[session]
	if !ok {
		return "", ErrMatchNotFound
	}
	key, keyOk := ma.newPlayerKey()
	if !keyOk {
		return "", ErrAllKeysGiven
	}
	return key, nil
}

func (m *Manager) PlayerInfo(key string) (*PlayerInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	session, ok := m.playerSessions[key]
	if !ok {
		return nil, ErrInvalidPlayerKey
	}
	ma := m.matches[session]
	return &PlayerInfo{
		PlayerID: int32(ma.checkPlayerKey(key)),
	}, nil
}

func (m *Manager) MatchInfo(session string) (*MatchInfo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	ma, ok := m.matches[session]
	if !ok {
		return nil, ErrMatchNotFound
	}
	info := ma.gameInfo
	return &MatchInfo{
		GameID:   info.ID,
		GameName: info.Name,
	}, nil
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

func newWebsocketConfig() centrifuge.WebsocketConfig {
	return centrifuge.WebsocketConfig{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (m *Manager) publishState(session string, state *game.State) error {
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("can't marshal state: %w", err)
	}
	_, err = m.node.Publish(session, data)
	return err
}

//nolint:gocritic // Centrifuge handler
func (m *Manager) onConnecting(
	ctx context.Context,
	event centrifuge.ConnectEvent,
) (centrifuge.ConnectReply, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.logger.Debug("connecting", zap.String("token", event.Token))
	if _, exists := m.playerSessions[event.Token]; !exists {
		return centrifuge.ConnectReply{}, centrifuge.DisconnectInvalidToken
	}
	return centrifuge.ConnectReply{
		Credentials: &centrifuge.Credentials{
			UserID: event.Token,
		},
	}, nil
}

func (m *Manager) onConnect(client *centrifuge.Client) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var (
		key     = client.UserID()
		session = m.playerSessions[client.UserID()]
		logger  = m.logger.With(
			zap.String("session", session),
			zap.String("key", key),
		)
	)
	ma, matchExists := m.matches[session]
	if !matchExists {
		client.Disconnect(centrifuge.DisconnectBadRequest)
		return
	}
	player := ma.checkPlayerKey(key)
	if player == 0 {
		client.Disconnect(centrifuge.DisconnectBadRequest)
		return
	}
	client.OnMessage(m.newMessageHandler(ma, session, player))
	client.OnSubscribe(func(se centrifuge.SubscribeEvent, sc centrifuge.SubscribeCallback) {
		if se.Channel != session {
			sc(centrifuge.SubscribeReply{}, centrifuge.ErrorBadRequest)
			return
		}
		sc(centrifuge.SubscribeReply{}, nil)
		if err := m.publishState(session, ma.game.State()); err != nil {
			logger.Debug("failed to publish state on player connect", zap.Error(err))
		}
	})
}

func (m *Manager) newMessageHandler(
	ma *match,
	session string,
	player int,
) centrifuge.MessageHandler {
	logger := m.logger.With(
		zap.String("session", session),
		zap.Int("player", player),
	)
	return func(me centrifuge.MessageEvent) {
		var message actionMessage
		if err := json.Unmarshal(me.Data, &message); err != nil {
			logger.Debug("message unmarshal failed", zap.Error(err))
			return
		}
		logger.Debug("received message", zap.Any("message", message))
		switch message.Type {
		case actionTypeClick:
			state, err := ma.game.HandleClick(&game.Click{
				Coords: game.Coords{
					X: message.X,
					Y: message.Y,
				},
				Player: player,
			})
			if err != nil {
				logger.Debug("handle click failed", zap.Error(err))
				return
			}
			if event := state.Event; event.IsGameEnd() {
				m.stopMatch(session)
			}
			if err := m.publishState(session, state); err != nil {
				m.logger.Debug("failed to publish game state", zap.Error(err))
			}
		case actionTypeMove:
			logger.Debug("received message with unsupported action type",
				zap.String("move", message.Type),
			)
		default:
			logger.Debug("received message with unknown action type",
				zap.String("action", message.Type),
			)
		}
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
		if _, exists := m.playerSessions[key]; !exists {
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
	info, err := m.storage.GameInfo(ctx, gameID)
	if err != nil {
		return "", fmt.Errorf("can't fetch game info: %w", err)
	}
	g, err := newGame(ctx, m.storage, gameID)
	if err != nil {
		return "", fmt.Errorf("can't create game: %w", err)
	}
	if _, err := g.Start(); err != nil {
		return "", fmt.Errorf("can't initialize game state: %w", err)
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
		gameInfo: info,
		game:     g,
		keys:     keys,
	}
	for _, key := range keys {
		m.playerSessions[key] = session
	}
	return session, nil
}

func (m *Manager) stopMatch(session string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	ma, exists := m.matches[session]
	if !exists {
		return
	}
	for _, key := range ma.keys {
		delete(m.playerSessions, key)
	}
	delete(m.matches, session)
}
