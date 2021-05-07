package matchesmanager

import (
	"context"
	"fmt"

	"github.com/mayamika/cellay/cellay-server/internal/gamesstorage"
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager/game"
	"github.com/mayamika/cellay/cellay-server/internal/token"
)

type match struct {
	gameID        int32
	game          *game.Game
	keysRequested int
	keys          []string
}

func newMatch(ctx context.Context, s *gamesstorage.Storage, gameID int32) (*match, error) {
	code, err := s.GameCode(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game code: %w", err)
	}
	assets, err := s.GameAssets(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("can't fetch game assets: %w", err)
	}
	layers := make([]string, 0)
	for name := range assets.Layers {
		layers = append(layers, name)
	}
	g, err := game.New(&game.Config{
		Code: code.Code,
		Field: game.Field{
			Cols: int(assets.Field.Cols),
			Rows: int(assets.Field.Rows),
		},
		Layers: layers,
	})
	if err != nil {
		return nil, fmt.Errorf("can't create game: %w", err)
	}
	keys, err := generateKeys(ctx, 2)
	if err != nil {
		return nil, fmt.Errorf("can't generate player keys: %w", err)
	}
	return &match{
		gameID: gameID,
		game:   g,
		keys:   keys,
	}, nil
}

func (m *match) newPlayerKey() (string, bool) {
	if m.keysRequested >= len(m.keys) {
		return "", false
	}
	key := m.keys[m.keysRequested]
	m.keysRequested++
	return key, true
}

//nolint:unused // Not implemented
func (m *match) checkPlayerKey(key string) int {
	for idx, playerKey := range m.keys {
		if key == playerKey {
			return idx + 1
		}
	}
	return 0
}

func generateKeys(ctx context.Context, n int) ([]string, error) {
	generatedKeys := make(map[string]struct{})
	for generated := 0; generated < n; generated++ {
		var key string
		for {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			key = token.New()
			isUnique := true
			for generatedKey := range generatedKeys {
				if key == generatedKey {
					isUnique = false
					break
				}
			}
			if isUnique {
				break
			}
		}
		generatedKeys[key] = struct{}{}
	}
	var keys []string
	for key := range generatedKeys {
		keys = append(keys, key)
	}
	return keys, nil
}
