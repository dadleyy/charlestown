package engine

import "log"
import "fmt"

const (
	statsMenuWidth  = 30
	statsMenuHeight = 4

	buildMenuHeight = 2
)

type uiRenderer struct {
	*log.Logger
	debug bool
}

func (renderer *uiRenderer) generate(state gameState) []renderable {
	texts := []string{fmt.Sprintf("mode: %d", state.cursor.mode), fmt.Sprintf("funds: %d", state.funds)}
	result := box(statsMenuWidth, statsMenuHeight, texts...)

	if renderer.debug {
		diagnostics := translate(box(state.dimensions.width-1, 2), point{0, state.dimensions.height - 3})
		result = append(result, diagnostics...)
	}

	if state.cursor.mode == cursorBuild {
		b := building{kind: state.cursor.building}
		text := fmt.Sprintf("building: %c (%s)", b.char(), &b)
		width := len(text)
		buildBox := box(width, buildMenuHeight, text)
		selection := translate(buildBox, point{state.dimensions.width - (width + 2), 0})
		result = append(result, selection...)
	}

	return result
}
