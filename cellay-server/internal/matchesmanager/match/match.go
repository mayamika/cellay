package match

import (
	"errors"
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

var (
	ErrCodeParseFailed = errors.New("code parse failed")
	ErrLuaCallFailed   = errors.New("lua function call failed")
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

const (
	matchStartFnName  = "init_game"
	handleMoveFnName  = "handle_move"
	handleClickFnName = "handle_click"
)

var (
	errMustBeFunction = errors.New("value must be a function")
)

type typeError struct {
	name   string
	reason error
}

func newTypeError(name string, reason error) *typeError {
	return &typeError{
		name:   name,
		reason: reason,
	}
}

func (te *typeError) Unwrap() error {
	return te.reason
}

func (te *typeError) Error() string {
	return fmt.Sprintf("value %q has invalid type: %s", te.name, te.reason)
}

func New(config *Config) (*Match, error) {
	m := &Match{
		lState: lua.NewState(),
		state: newState(
			config.Field.Cols,
			config.Field.Rows,
			config.Layers,
		),
	}
	if err := m.parseCode(config.Code); err != nil {
		return nil, wrapErr(ErrCodeParseFailed, err)
	}
	return m, nil
}

func (m *Match) Start() (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.callStart(); err != nil {
		return nil, wrapErr(ErrLuaCallFailed, err)
	}
	return stateCopy(m.state), nil
}

func (m *Match) HandleClick(click *Click) (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.callHandleClick(click); err != nil {
		return nil, wrapErr(ErrLuaCallFailed, err)
	}
	return stateCopy(m.state), nil
}

func (m *Match) HandleMove(move *Move) (*State, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return stateCopy(m.state), errors.New("not implemented") //nolint:goerr113 // Not implemented
}

func (m *Match) parseCode(code string) error {
	if err := m.lState.DoString(code); err != nil {
		return err
	}
	registerTypes(m.lState)
	var err error
	m.startFn, err = loadLuaFunction(m.lState, matchStartFnName, 0)
	if err != nil {
		return err
	}
	m.handleMoveFn, err = loadLuaFunction(m.lState, handleMoveFnName, 0)
	if err != nil {
		return err
	}
	m.handleClickFn, err = loadLuaFunction(m.lState, handleClickFnName, 0)
	if err != nil {
		return err
	}
	return nil
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
		return lua.P{}, newTypeError(name, errMustBeFunction)
	}
	return lua.P{
		Fn:      fn,
		NRet:    nRet,
		Protect: true,
	}, nil
}

func wrapErr(err, reason error) error {
	return fmt.Errorf("%w: %v", err, reason)
}
