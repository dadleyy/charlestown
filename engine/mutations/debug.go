package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Debug toggles the debug flag on the game.
func Debug() Mutation {
	return debug{}
}

type debug struct {
}

func (d debug) Apply(game objects.Game) objects.Game {
	next := game
	next.Debug = !next.Debug
	return next
}
