package match

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	t.Run("TicTacToe", testTicTacToe)
}

func testTicTacToe(t *testing.T) {
	code := `
	turn = 1

	function passTurn()
		turn = 3 - turn
	end

	function init_game(state)
		-- game have no initial state
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
		passTurn()
	end
	`
	r := require.New(t)
	match, err := New(&Config{
		Code: code,
		Field: GameField{
			Cols: 3,
			Rows: 3,
		},
		Layers: []string{"main"},
	})
	r.NoError(err)
	/*
		___
		___
		___
	*/
	state, err := match.Start()
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
	/*
		___
		_x_
		___
	*/
	testCorrectTicTacToeMove(t, match, state, 1, 1, 1)
	testOutOfBoundsTicTacToeMove(t, match, 2, 5, 4)
	testAnotherPlayersTurnTicTacToeMove(t, match, state, 1, 2, 0)
	/*
		__0
		_x_
		___
	*/
	testCorrectTicTacToeMove(t, match, state, 2, 2, 0)
	/*
		__0
		_x_
		__x
	*/
	testCorrectTicTacToeMove(t, match, state, 1, 2, 2)
	/*
		0_0
		_x_
		__x
	*/
	testCorrectTicTacToeMove(t, match, state, 2, 0, 0)
	/*
		0x0
		_x_
		__x
	*/
	testCorrectTicTacToeMove(t, match, state, 1, 1, 0)
	/*
		0x0
		0x_
		__x
	*/
	testCorrectTicTacToeMove(t, match, state, 2, 0, 1)
	/*
		0x0
		0x_
		_xx
	*/
	testCorrectTicTacToeMove(t, match, state, 1, 1, 2)
}

func testCorrectTicTacToeMove(t *testing.T, m *Match, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	state.Table["main"][x][y] = player
	newState, err := m.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}

func testOutOfBoundsTicTacToeMove(t *testing.T, m *Match, player, x, y int) {
	t.Helper()
	r := require.New(t)
	newState, err := m.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.Error(err)
	r.Nil(newState)
}

func testAnotherPlayersTurnTicTacToeMove(t *testing.T, m *Match, state *State, player, x, y int) {
	t.Helper()
	r := require.New(t)
	newState, err := m.HandleClick(&Click{
		Coords: Coords{
			X: x,
			Y: y,
		},
		Player: player,
	})
	r.NoError(err)
	r.Equal(state, newState)
}
