package game

import lua "github.com/yuin/gopher-lua"

//go:generate ${TOOLS_PATH}/enumer -type EventType -trimprefix EventType -json -text -transform=snake
type EventType int8

const (
	EventTypeUnknown EventType = iota
	EventTypeWin
	EventTypeDraw
	EventTypeNotification
)

type Event struct {
	Type    EventType
	Player  int    `json:",omitempty"`
	Message string `json:",omitempty"`
}

const (
	luaEventTypename = "Event"
)

func registerEventType(lState *lua.LState) {
	methods := map[string]lua.LGFunction{}
	mt := lState.NewTypeMetatable(luaEventTypename)
	lState.SetGlobal(luaEventTypename, mt)
	lState.SetField(mt, "win", lState.NewFunction(newLuaWinEvent))
	lState.SetField(mt, "draw", lState.NewFunction(newLuaDrawEvent))
	lState.SetField(mt, "notification", lState.NewFunction(newLuaNotificationEvent))
	lState.SetField(mt, "__index", lState.SetFuncs(lState.NewTable(), methods))
}

func pushEvent(lState *lua.LState, ev *Event) {
	ud := lState.NewUserData()
	ud.Value = ev
	lState.SetMetatable(ud, lState.GetTypeMetatable(luaEventTypename))
	lState.Push(ud)
}

func newLuaWinEvent(lState *lua.LState) int {
	checkArgsCount(lState, 1)
	pushEvent(lState, &Event{
		Type:   EventTypeWin,
		Player: lState.CheckInt(1),
	})
	return 1
}

func newLuaDrawEvent(lState *lua.LState) int {
	checkArgsCount(lState, 0)
	pushEvent(lState, &Event{
		Type: EventTypeDraw,
	})
	return 1
}

func newLuaNotificationEvent(lState *lua.LState) int {
	checkArgsCount(lState, 1)
	pushEvent(lState, &Event{
		Type:    EventTypeNotification,
		Message: lState.CheckString(1),
	})
	return 1
}

func eventFromLua(lState *lua.LState, n int) *Event {
	if ud, ok := lState.Get(n).(*lua.LUserData); ok {
		if val := ud.Value.(*Event); ok {
			return val
		}
	}
	return nil
}
