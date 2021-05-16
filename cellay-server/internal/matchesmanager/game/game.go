package game

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

type Game struct {
	mu            sync.Mutex
	lState        *lua.LState
	startFn       lua.P
	handleMoveFn  lua.P
	handleClickFn lua.P
	state         *State
}

type Field struct {
	Cols, Rows int
}

type Config struct {
	Code   string
	Field  Field
	Layers []string
}

type Coords struct {
	X int
	Y int
}

const (
	gameStartFnName   = "init_game"
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

func New(config *Config) (*Game, error) {
	g := &Game{
		lState: lua.NewState(),
		state: newState(
			config.Field.Cols,
			config.Field.Rows,
			config.Layers,
		),
	}
	if err := g.parseCode(config.Code); err != nil {
		return nil, wrapErr(ErrCodeParseFailed, err)
	}
	return g, nil
}

func (g *Game) Start() (*State, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if err := g.callStart(); err != nil {
		return nil, wrapErr(ErrLuaCallFailed, err)
	}
	return stateCopy(g.state), nil
}

func (g *Game) HandleClick(click *Click) (*State, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if err := g.callHandleClick(click); err != nil {
		return nil, wrapErr(ErrLuaCallFailed, err)
	}
	return stateCopy(g.state), nil
}

func (g *Game) State() *State {
	g.mu.Lock()
	defer g.mu.Unlock()
	return stateCopy(g.state)
}

func (g *Game) HandleMove(move *Move) (*State, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return stateCopy(g.state), errors.New("not implemented") //nolint:goerr113 // Not implemented
}

func (g *Game) parseCode(code string) error {
	if err := g.lState.DoString(code); err != nil {
		return err
	}
	registerTypes(g.lState)
	var err error
	g.startFn, err = loadLuaFunction(g.lState, gameStartFnName, 0)
	if err != nil {
		return err
	}
	g.handleMoveFn, err = loadLuaFunction(g.lState, handleMoveFnName, 1)
	if err != nil {
		return err
	}
	g.handleClickFn, err = loadLuaFunction(g.lState, handleClickFnName, 1)
	if err != nil {
		return err
	}
	return nil
}

func registerTypes(lState *lua.LState) {
	registerClickType(lState)
	registerEventType(lState)
	registerStateType(lState)
}

func (g *Game) callStart() error {
	return g.lState.CallByParam(
		g.startFn,
		newStateUserData(g.lState, g.state),
	)
}

func (g *Game) callHandleClick(click *Click) error {
	if err := g.lState.CallByParam(
		g.handleClickFn,
		newStateUserData(g.lState, g.state),
		newClickUserData(g.lState, click),
	); err != nil {
		return err
	}
	g.state.Event = eventFromLua(g.lState, 1)
	g.lState.Pop(1)
	return nil
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

func checkArgsCount(lState *lua.LState, argsRequired int) {
	if lState.GetTop() != argsRequired {
		lState.ArgError(argsRequired, fmt.Sprintf("%d args expected", argsRequired))
	}
}

func wrapErr(err, reason error) error {
	return fmt.Errorf("%w: %v", err, reason)
}
