package engine

import "testing"
import "github.com/franela/goblin"

func TestMutations(t *testing.T) {
	g := goblin.Goblin(t)

	var mut mutation

	g.Describe("move", func() {
		g.It("should prevent the movement to a position beyond boundaries", func() {
			mut = move(0, 1)
			g.Assert(mut(gameState{}).cursor.location.x).Equal(0)
			g.Assert(mut(gameState{}).cursor.location.y).Equal(0)
		})

		g.It("should prevent a negative move", func() {
			world := dimensions{width: 10, height: 10}
			mut = move(-1, -1)
			g.Assert(mut(gameState{world: world}).cursor.location.x).Equal(1)
			g.Assert(mut(gameState{world: world}).cursor.location.y).Equal(1)
		})

		g.It("should move the cursor where available", func() {
			world := dimensions{width: 10, height: 10}
			mut = move(1, 1)
			g.Assert(mut(gameState{world: world}).cursor.location.x).Equal(1)
			g.Assert(mut(gameState{world: world}).cursor.location.y).Equal(1)
		})
	})
}
