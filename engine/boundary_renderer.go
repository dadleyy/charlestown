package engine

import "log"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/drawing"

type boundaryRenderer struct {
	*log.Logger
}

func (renderer boundaryRenderer) generate(state objects.Game) []objects.Renderable {
	bounding := drawing.Box(state.World.Width, state.World.Height)

	midx, midy := state.Dimensions.Midway()
	point := objects.Point{
		X: midx - state.Cursor.Location.X,
		Y: midy - state.Cursor.Location.Y,
	}

	return drawing.Translate(bounding, point)
}
