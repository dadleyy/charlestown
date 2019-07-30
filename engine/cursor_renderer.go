package engine

import "log"

type cursorRenderer struct {
	*log.Logger
}

func (renderer *cursorRenderer) generate(state gameState) []renderable {
	x, y := state.dimensions.midway()
	return []renderable{renderable{location: point{x, y}, symbol: state.cursor.char()}}
}
