package match

import (
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

type Match struct {
	mu     sync.Mutex
	lState *lua.LState
}

type Move struct{}

const (
	matchStartFuncName = "game_init"
	handleMoveFuncName = "handle_move"
)

func New(code string) (*Match, error) {
	lState := lua.NewState()
	if err := lState.DoString(code); err != nil {
		return nil, fmt.Errorf("can't run code: %w", err)
	}
	initFn := lState.GetGlobal(matchStartFuncName)
	if initFn.Type() != lua.LTFunction {
		return nil, fmt.Errorf("game_init must be a function")
	}
	return &Match{
		lState: lState,
	}, nil
}

func (m *Match) Start() ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return nil, nil
}

func (m *Match) HandleMove(move *Move) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return nil, nil
}
