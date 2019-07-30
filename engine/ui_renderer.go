package engine

import "log"
import "fmt"

const (
	statsMenuWidth  = 30
	statsMenuHeight = 4

	buildMenuWidth  = 20
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
		buildBox := box(buildMenuWidth, buildMenuHeight, fmt.Sprintf("building: %c", b.char()))
		selection := translate(buildBox, point{state.dimensions.width - (buildMenuWidth + 2), 0})
		result = append(result, selection...)
	}

	return result
}
