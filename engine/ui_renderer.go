package engine

import "log"
import "fmt"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/drawing"
import "github.com/dadleyy/charlestown/engine/constants"

const (
	statsMenuWidth  = 30
	statsMenuHeight = 4

	buildMenuHeight = 2
)

type uiRenderer struct {
	*log.Logger
	debug bool
}

func (renderer *uiRenderer) generate(state objects.Game) []objects.Renderable {
	texts := []string{
		fmt.Sprintf("mode: %d", state.Cursor.Mode),
		fmt.Sprintf("funds: %d (+%d)", state.Funds, state.Revenue),
	}
	result := drawing.Box(statsMenuWidth, statsMenuHeight, texts...)

	if renderer.debug {
		w := state.Dimensions.Width - 1
		h := state.Dimensions.Height - 3
		t := fmt.Sprintf("player location: %s (%d buildings) (frame %d)", &state.Cursor, len(state.Buildings), state.Frame)
		box := drawing.Box(w, 2, t)
		diagnostics := drawing.Translate(box, objects.Point{0, h})
		result = append(result, diagnostics...)
	}

	/*
		for index, message := range state.Messages {
			top := state.Dimensions.Height - 10 - (index * 4)
			renderer.Printf("rendering message #%d '%s'", index, message.Text)
			container := drawing.Translate(drawing.Box(len(message.Text)+1, 3, message.Text), objects.Point{0, top})
			result = append(result, container...)
		}
	*/

	if state.Cursor.Mode == constants.CursorBuild {
		b := objects.Building{Kind: state.Cursor.Building}
		text := fmt.Sprintf("building: %c (%s)", b.Char(), &b)
		width := len(text)
		buildBox := drawing.Box(width, buildMenuHeight, text)
		selection := drawing.Translate(buildBox, objects.Point{state.Dimensions.Width - (width + 2), 0})
		result = append(result, selection...)
	}

	return result
}
