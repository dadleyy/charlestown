package mutations

import "github.com/dadleyy/charlestown/engine/objects"

func dup(state objects.Game) objects.Game {
	return objects.Game{
		Revenue:    state.Revenue,
		Funds:      state.Funds,
		Buildings:  state.Buildings[0:],
		World:      state.World.Dup(),
		Dimensions: state.Dimensions.Dup(),
		Cursor:     state.Cursor.Dup(),
		Flashes:    state.Flashes[0:],
		Frame:      state.Frame,
	}
}
