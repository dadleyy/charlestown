package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Resize mutations update the game's viewport dimensions.
func Resize(width, height int) Mutation {
	return resize{width, height}
}

type resize struct {
	width  int
	height int
}

func (r resize) Apply(game objects.Game) objects.Game {
	next := game
	next.Dimensions = objects.Dimensions{r.width, r.height}
	return next
}
