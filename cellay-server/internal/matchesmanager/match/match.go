package match

import (
	"errors"
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

type Match struct {
	mu            sync.Mutex
	lState        *lua.LState
	startFn       lua.P
	handleMoveFn  lua.P
	handleClickFn lua.P
	state         *State
}

type GameField struct {
	Cols, Rows int
}

type Config struct {
	Code   string
	Field  GameField
	Layers []string
}

type Coords struct {
	X int
	Y int
}

type Move struct {
	From   Coords
	To     Coords
	Player int
}

const (
	matchStartFuncName  = "init_game"
	handleMoveFuncName  = "handle_move"
	handleClickFuncName = "handle_click"
)

func New(config *Config) (*Match, error) {
	lState := lua.NewState()
	if err := lState.DoString(config.Code); err != nil {
		return nil, fmt.Errorf("can't run code: %w", err)
	}
	registerTypes(lState)
	startFn, err := loadLuaFunction(lState, matchStartFuncName, 1)
	if err != nil {
		return nil, fmt.Errorf("can't load start fn: %w", err)
	}
	handleMoveFn, err := loadLuaFunction(lState, handleMoveFuncName, 1)
	if err != nil {
		return nil, fmt.Errorf("can't load handle move fn: %w", err)
	}
	handleClickFn, err := loadLuaFunction(lState, handleClickFuncName, 1)
	if err != nil {
		return nil, fmt.Errorf("can't load handle click fn: %w", err)
	}
	return &Match{
		lState:        lState,
		startFn:       startFn,
		handleMoveFn:  handleMoveFn,
		handleClickFn: handleClickFn,
		state: newState(
			config.Field.Cols,
			config.Field.Rows,
			config.Layers,
		),
	}, nil
}

func (m *Match) Start() (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.callStart(); err != nil {
		return nil, err
	}
	return stateCopy(m.state), nil
}

func (m *Match) HandleClick(click *Click) (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.callHandleClick(click); err != nil {
		return nil, err
	}
	return stateCopy(m.state), nil
}

func (m *Match) HandleMove(move *Move) (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return stateCopy(m.state), errors.New("not implemented")
}

func registerTypes(lState *lua.LState) {
	registerClickType(lState)
	registerStateType(lState)
}

func (m *Match) callStart() error {
	return m.lState.CallByParam(
		m.startFn,
		newStateUserData(m.lState, m.state),
	)
}

func (m *Match) callHandleClick(click *Click) error {
	return m.lState.CallByParam(
		m.handleClickFn,
		newStateUserData(m.lState, m.state),
		newClickUserData(m.lState, click),
	)
}

func loadLuaFunction(lState *lua.LState, name string, nRet int) (lua.P, error) {
	fn := lState.GetGlobal(name)
	if fn.Type() != lua.LTFunction {
		return lua.P{}, errors.New("value must be a function")
	}
	return lua.P{
		Fn:      fn,
		NRet:    nRet,
		Protect: true,
	}, nil
}
