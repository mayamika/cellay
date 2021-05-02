package match

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	t.Run("SimpleCode", testSimpleCode)
}

func testSimpleCode(t *testing.T) {
	code := `
	function game_init()
	end

	function handle_move()
	end
	`
	r := require.New(t)
	_, err := New(code)
	r.NoError(err)
}
