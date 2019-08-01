package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Resize mutations update the game's viewport dimensions.
func Resize(width, height int) Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)
		next.Dimensions = objects.Dimensions{width, height}
		return next
	}
}
