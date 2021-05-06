package game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	t.Run("TicTacToe", testTicTacToe)
}

func testTicTacToe(t *testing.T) {
	code := `
	turn = 1

	function passTurn()
		turn = 3 - turn
	end

	function check_win_vertical(state, player)
		for x = 0, 2 do
			ok = true
			for y = 1, 2 do
				lhs, rhs = state:fieldTile(x, y, "main"), state:fieldTile(x, y - 1, "main")
				equal = (lhs == player) and (lhs == rhs)
				ok = ok and equal
			end
			if ok then
				return true
			end
		end
		return false
	end

	function check_win_horizontal(state, player)
		for y = 0, 2 do
			ok = true
			for x = 1, 2 do
				lhs, rhs = state:fieldTile(x, y, "main"), state:fieldTile(x - 1, y, "main")
				equal = (lhs == player) and (lhs == rhs)
				ok = ok and equal
			end
			if ok then
				return true
			end
		end
		return false
	end

	function check_win_diagonal(state, player)
		ok = true
		for n = 1, 2 do
			lhs, rhs = state:fieldTile(n, n, "main"), state:fieldTile(n - 1, n - 1, "main")
			ok = ok and (lhs == player and lhs == rhs)
		end
		if ok then
			return true
		end
		ok = true
		for n = 1, 2 do
			lhs, rhs = state:fieldTile(2 - n, n, "main"), state:fieldTile(3 - n, n - 1, "main")
			ok = ok and (lhs == player and lhs == rhs)
		end
		return ok
	end

	function check_win(state, player)
		if check_win_vertical(state, player) then
			return true
		end
		if check_win_horizontal(state, player) then
			return true
		end
		if check_win_diagonal(state, player) then
			return true
		end
		return false
	end

	-- check_draw function always called after win check,
	-- so we can simplify the check
	function check_draw(state)
		for x = 0, 2 do
			for y = 0, 2 do
				if state:fieldTile(x, y, "main") == 0 then
					return false
				end
			end
		end
		return true
	end

	function init_game(state)
		-- game has no initial state
	end

	function handle_move(state, move)
		-- game doesn't support moves, only clicks supported
	end

	function handle_click(state, click)
		if turn ~= click:player() then
			-- another player's turn, ignore click
			return
		end
		if state:fieldTile(click:x(), click:y(), "main") ~= 0 then
			-- field already occupied, ignore click
			return
		end
		state:setFieldTile(click:x(), click:y(), "main", click:player())
		if check_win(state, click:player()) then
			return Event.win(click:player())
		end
		if check_draw(state) then
			return Event.draw()
		end
		passTurn()
	end
	`
	t.Run("VerticalWin", func(t *testing.T) { testTicTacToeVerticalWin(t, code) })
	t.Run("HorizontalWin", func(t *testing.T) { testTicTacToeHorizontalWin(t, code) })
	t.Run("DiagonalWin", func(t *testing.T) { testTicTacToeDiagonalWin(t, code) })
	t.Run("Draw", func(t *testing.T) { testTicTacToeDraw(t, code) })
}

func testTicTacToeVerticalWin(t *testing.T, code string) {
	r := require.New(t)
	game, err := New(&Config{
		Code: code,
		Field: Field{
			Cols: 3,
			Rows: 3,
		},
		Layers: []string{"main"},
	})
	r.NoError(err)
	// ___
	// ___
	// ___
	state, err := game.Start()
	r.NoError(err)
	r.Equal(&State{
		Table: map[string][][]int{
			"main": {
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}, state)
	// ___
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 1, 1, 1)
	testOutOfBoundsTicTacToeMove(t, game, 2, 5, 4)
	testAnotherPlayersTurnTicTacToeMove(t, game, state, 1, 2, 0)
	// __0
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 2, 2, 0)
	// __0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 1, 2, 2)
	// 0_0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 2, 0, 0)
	// 0x0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 1, 1, 0)
	// 0x0
	// 0x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 2, 0, 1)
	// 0x0
	// 0x_
	// _xx
	testWinTicTacToeMove(t, game, state, 1, 1, 2)
}

func testTicTacToeHorizontalWin(t *testing.T, code string) {
	r := require.New(t)
	game, err := New(&Config{
		Code: code,
		Field: Field{
			Cols: 3,
			Rows: 3,
		},
		Layers: []string{"main"},
	})
	r.NoError(err)
	// ___
	// ___
	// ___
	state, err := game.Start()
	r.NoError(err)
	r.Equal(&State{
		Table: map[string][][]int{
			"main": {
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}, state)
	// ___
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 1, 1, 1)
	testOutOfBoundsTicTacToeMove(t, game, 2, 5, 4)
	testAnotherPlayersTurnTicTacToeMove(t, game, state, 1, 2, 0)
	// __0
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 2, 2, 0)
	// __0
	// _xx
	// ___
	testCorrectTicTacToeMove(t, game, state, 1, 2, 1)
	// _00
	// _xx
	// ___
	testCorrectTicTacToeMove(t, game, state, 2, 1, 0)
	// _00
	// xxx
	// ___
	testWinTicTacToeMove(t, game, state, 1, 0, 1)
}

func testTicTacToeDiagonalWin(t *testing.T, code string) {
	r := require.New(t)
	game, err := New(&Config{
		Code: code,
		Field: Field{
			Cols: 3,
			Rows: 3,
		},
		Layers: []string{"main"},
	})
	r.NoError(err)
	// ___
	// ___
	// ___
	state, err := game.Start()
	r.NoError(err)
	r.Equal(&State{
		Table: map[string][][]int{
			"main": {
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}, state)
	// ___
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 1, 1, 1)
	testOutOfBoundsTicTacToeMove(t, game, 2, 5, 4)
	testAnotherPlayersTurnTicTacToeMove(t, game, state, 1, 2, 0)
	// __0
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 2, 2, 0)
	// __0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 1, 2, 2)
	// _00
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 2, 1, 0)
	// x00
	// _x_
	// __x
	testWinTicTacToeMove(t, game, state, 1, 0, 0)
}

func testTicTacToeDraw(t *testing.T, code string) {
	r := require.New(t)
	game, err := New(&Config{
		Code: code,
		Field: Field{
			Cols: 3,
			Rows: 3,
		},
		Layers: []string{"main"},
	})
	r.NoError(err)
	// ___
	// ___
	// ___
	state, err := game.Start()
	r.NoError(err)
	r.Equal(&State{
		Table: map[string][][]int{
			"main": {
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}, state)
	// ___
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 1, 1, 1)
	testOutOfBoundsTicTacToeMove(t, game, 2, 5, 4)
	testAnotherPlayersTurnTicTacToeMove(t, game, state, 1, 2, 0)
	// __0
	// _x_
	// ___
	testCorrectTicTacToeMove(t, game, state, 2, 2, 0)
	testOutOfBoundsTicTacToeMove(t, game, 1, 2, 3)
	testAnotherPlayersTurnTicTacToeMove(t, game, state, 2, 0, 1)
	// __0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 1, 2, 2)
	// 0_0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 2, 0, 0)
	// 0x0
	// _x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 1, 1, 0)
	// 0x0
	// 0x_
	// __x
	testCorrectTicTacToeMove(t, game, state, 2, 0, 1)
	// 0x0
	// 0x_
	// x_x
	testCorrectTicTacToeMove(t, game, state, 1, 0, 2)
	// 0x0
	// 0x_
	// x0x
	testCorrectTicTacToeMove(t, game, state, 2, 1, 2)
	// 0x0
	// 0xx
	// x0x
	testDrawTicTacToeMove(t, game, state, 1, 2, 1)
}

func testCorrectTicTacToeMove(t *testing.T, g *Game, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	state.Table["main"][x][y] = player
	newState, err := g.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}

func testWinTicTacToeMove(t *testing.T, g *Game, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	state.Table["main"][x][y] = player
	state.Event = &Event{
		Type:   EventTypeWin,
		Player: player,
	}
	newState, err := g.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}

func testDrawTicTacToeMove(t *testing.T, g *Game, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	state.Table["main"][x][y] = player
	state.Event = &Event{
		Type: EventTypeDraw,
	}
	newState, err := g.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}

func testOutOfBoundsTicTacToeMove(t *testing.T, g *Game, player, x, y int) {
	t.Helper()
	r := require.New(t)
	newState, err := g.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.Error(err)
	r.Nil(newState)
}

func testAnotherPlayersTurnTicTacToeMove(t *testing.T, g *Game, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	newState, err := g.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}
