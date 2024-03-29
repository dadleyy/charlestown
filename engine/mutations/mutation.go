package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Mutation functions are applied to a given game state, returning an updated state.
type Mutation interface {
	Apply(objects.Game) objects.Game
}
