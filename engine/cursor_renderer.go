package engine

import "log"
import "image/color"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

type cursorRenderer struct {
	*log.Logger
}

func (renderer *cursorRenderer) hitColor(mode int) color.Color {
	switch mode {
	case constants.CursorDemolish:
		return color.RGBA{255, 0, 0, 1.0}
	case constants.CursorMove:
		return color.RGBA{0, 255, 255, 1.0}
	default:
		return color.RGBA{255, 255, 255, 1.0}
	}
}

func (renderer *cursorRenderer) generate(state objects.Game) []objects.Renderable {
	cursor := objects.Renderable{state.Dimensions.MidwayPoint(), state.Cursor.Char(), nil}

	for _, b := range state.Buildings {
		hit := b.Location.Equals(state.Cursor.Location)

		if hit {
			cursor.Color = renderer.hitColor(state.Cursor.Mode)
		}
	}

	return []objects.Renderable{cursor}
}
