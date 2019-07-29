package engine

import "log"

type cursorRenderer struct {
	*log.Logger
}

func (renderer *cursorRenderer) generate(state gameState) []renderable {
	x := state.dimensions.width / 2
	y := state.dimensions.height / 2
	return []renderable{renderable{location: point{x, y}, symbol: state.cursor.char()}}
}
