package matchesmanager

import (
	"github.com/mayamika/cellay/cellay-server/internal/matchesmanager/game"
)

type match struct {
	gameID        int32
	game          *game.Game
	keysRequested int
	keys          []string
}

func (m *match) newPlayerKey() (string, bool) {
	if m.keysRequested >= len(m.keys) {
		return "", false
	}
	key := m.keys[m.keysRequested]
	m.keysRequested++
	return key, true
}

func (m *match) checkPlayerKey(key string) int {
	for idx, playerKey := range m.keys {
		if key == playerKey {
			return idx + 1
		}
	}
	return 0
}
