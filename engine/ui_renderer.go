package engine

import "log"
import "fmt"

type uiRenderer struct {
	*log.Logger
	debug bool
}

func (renderer *uiRenderer) generate(state gameState) []renderable {
	result := box(30, 4, fmt.Sprintf("mode: %d", state.cursor.mode))

	if renderer.debug {
		diagnostics := translate(box(state.dimensions.width-1, 2), point{0, state.dimensions.height - 3})
		result = append(result, diagnostics...)
	}

	return result
}
