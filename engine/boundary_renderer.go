package engine

import "log"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/drawing"

type boundaryRenderer struct {
	*log.Logger
}

func (renderer boundaryRenderer) generate(state objects.Game) []objects.Renderable {
	bounding := drawing.Box(state.World.Width, state.World.Height-1)
	mid := state.Dimensions.MidwayPoint()
	return drawing.Translate(bounding, mid.Translate(-state.Cursor.Location.X, -state.Cursor.Location.Y))
}
