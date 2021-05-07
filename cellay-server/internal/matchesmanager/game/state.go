package game

import (
	lua "github.com/yuin/gopher-lua"
)

type State struct {
	Table map[string][][]int // maps layer name to [col][row]
	Event *Event             `json:",omitempty"`
}

const (
	luaStateTypename = "State"
)

func newState(cols, rows int, layers []string) *State {
	s := &State{
		Table: make(map[string][][]int),
	}
	for _, layer := range layers {
		s.Table[layer] = make([][]int, cols)
		for col := range s.Table[layer] {
			s.Table[layer][col] = make([]int, rows)
		}
	}
	return s
}

func stateCopy(s *State) *State {
	sCopy := &State{
		Table: make(map[string][][]int),
	}
	if ev := s.Event; ev != nil {
		sCopy.Event = &Event{}
		*sCopy.Event = *ev
	}
	for layer, layerTable := range s.Table {
		sCopy.Table[layer] = make([][]int, len(layerTable))
		for col, rows := range layerTable {
			sCopy.Table[layer][col] = make([]int, len(rows))
			copy(sCopy.Table[layer][col], s.Table[layer][col])
		}
	}
	return sCopy
}

func newStateUserData(lState *lua.LState, state *State) *lua.LUserData {
	ud := lState.NewUserData()
	ud.Value = state
	lState.SetMetatable(ud, lState.GetTypeMetatable(luaStateTypename))
	return ud
}

func registerStateType(lState *lua.LState) {
	methods := map[string]lua.LGFunction{
		"setFieldTile": luaStateSetFieldTile,
		"fieldTile":    luaStateFieldTile,
	}
	mt := lState.NewTypeMetatable(luaStateTypename)
	lState.SetGlobal(luaStateTypename, mt)
	lState.SetField(mt, "__index", lState.SetFuncs(lState.NewTable(), methods))
}

func checkLuaState(lState *lua.LState, n int) *State {
	ud := lState.CheckUserData(n)
	if val, ok := ud.Value.(*State); ok {
		return val
	}
	lState.ArgError(n, "state expected")
	return nil
}

func luaStateSetFieldTile(lState *lua.LState) int {
	checkArgsCount(lState, 5)
	var (
		state = checkLuaState(lState, 1)
		x     = lState.CheckInt(2)
		y     = lState.CheckInt(3)
		layer = lState.CheckString(4)
		tile  = lState.CheckInt(5)
	)
	layerTable, ok := state.Table[layer]
	if !ok {
		lState.ArgError(4, "unexpected asset layer")
	}
	if x >= len(layerTable) {
		lState.ArgError(2, "x is out of bounds")
	}
	if y >= len(layerTable[x]) {
		lState.ArgError(3, "y is out of bounds")
	}
	layerTable[x][y] = tile
	return 0
}

func luaStateFieldTile(lState *lua.LState) int {
	checkArgsCount(lState, 4)
	var (
		state = checkLuaState(lState, 1)
		x     = lState.CheckInt(2)
		y     = lState.CheckInt(3)
		layer = lState.CheckString(4)
	)
	layerTable, ok := state.Table[layer]
	if !ok {
		lState.ArgError(4, "unexpected asset layer")
	}
	if x >= len(layerTable) {
		lState.ArgError(2, "x is out of bounds")
	}
	if y >= len(layerTable[x]) {
		lState.ArgError(3, "y is out of bounds")
	}
	lState.Push(lua.LNumber(layerTable[x][y]))
	return 1
}
