package engine

import "log"

type cursorRenderer struct {
	*log.Logger
}

func (renderer *cursorRenderer) generate(state gameState) []renderable {
	x, y := state.dimensions.midway()
	renderer.Printf("rendering cursor for mode %b", state.cursor.building)
	return []renderable{renderable{location: point{x, y}, symbol: state.cursor.char()}}
}
