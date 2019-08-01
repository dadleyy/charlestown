package engine

import "log"
import "github.com/dadleyy/charlestown/engine/objects"

type cursorRenderer struct {
	*log.Logger
}

func (renderer *cursorRenderer) generate(state objects.Game) []objects.Renderable {
	x, y := state.Dimensions.Midway()
	cursor := objects.Renderable{Location: objects.Point{x, y}, Symbol: state.Cursor.Char()}
	return []objects.Renderable{cursor}
}
