package match

import lua "github.com/yuin/gopher-lua"

type Click struct {
	Coords
	Player int
}

const (
	luaClickTypename = "Click"
)

func newClickUserData(lState *lua.LState, click *Click) *lua.LUserData {
	ud := lState.NewUserData()
	ud.Value = click
	lState.SetMetatable(ud, lState.GetTypeMetatable(luaClickTypename))
	return ud
}

func registerClickType(lState *lua.LState) {
	methods := map[string]lua.LGFunction{
		"x":      luaClickX,
		"y":      luaClickY,
		"player": luaClickPlayer,
	}
	mt := lState.NewTypeMetatable(luaClickTypename)
	lState.SetGlobal(luaClickTypename, mt)
	lState.SetField(mt, "__index", lState.SetFuncs(lState.NewTable(), methods))
}

func checkLuaClick(lState *lua.LState) *Click {
	ud := lState.CheckUserData(1)
	if val, ok := ud.Value.(*Click); ok {
		return val
	}
	lState.ArgError(1, "click expected")
	return nil
}

func luaClickX(lState *lua.LState) int {
	click := checkLuaClick(lState)
	lState.Push(lua.LNumber(click.X))
	return 1
}

func luaClickY(lState *lua.LState) int {
	click := checkLuaClick(lState)
	lState.Push(lua.LNumber(click.Y))
	return 1
}

func luaClickPlayer(lState *lua.LState) int {
	click := checkLuaClick(lState)
	lState.Push(lua.LNumber(click.Player))
	return 1
}
