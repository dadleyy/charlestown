package engine

import "log"

type uiRenderer struct {
	*log.Logger
}

func (renderer *uiRenderer) generate(state gameState) []renderable {
	result := box(30, 4)
	return result
}
