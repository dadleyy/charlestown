package mutations

import "fmt"
import "testing"
import "encoding/json"
import "github.com/franela/goblin"
import "github.com/dadleyy/charlestown/engine/objects"

func TestDebug(t *testing.T) {
	g := goblin.Goblin(t)

	subject := func(actual objects.Game, expected objects.Game) (string, string, error) {
		a, e := json.Marshal(actual)

		if e != nil {
			return "", "", e
		}
		exp, e := json.Marshal(expected)
		return fmt.Sprintf("%s", a), fmt.Sprintf("%s", exp), e
	}

	g.Describe("Debug mutation", func() {
		g.It("returns a copy of the game with the debug flag flipped", func() {
			game := objects.Game{}
			actual, expected, e := subject(Debug().Apply(game), objects.Game{Debug: true})
			g.Assert(e).Equal(nil)
			g.Assert(actual).Equal(expected)
		})
	})
}
