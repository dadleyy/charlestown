package engine

import "log"

type boundaryRenderer struct {
	*log.Logger
}

func (renderer boundaryRenderer) generate(state gameState) []renderable {
	bounding := box(state.world.width, state.world.height)
	midx, midy := state.dimensions.midway()
	return translate(bounding, point{midx - state.cursor.location.x, midy - state.cursor.location.y})
}
