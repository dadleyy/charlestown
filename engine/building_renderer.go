package engine

import "log"
import "image/color"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

type buildingRenderer struct {
	*log.Logger
}

func (renderer *buildingRenderer) generate(state objects.Game) []objects.Renderable {
	buildings := make([]objects.Renderable, 0, len(state.Buildings))

	hx, hy := state.Dimensions.Midway()
	cx, cy := state.Cursor.Location.Values()

	for _, building := range state.Buildings {
		x, y := building.Location.Values()
		projected := objects.Point{x + hx - cx, y + hy - cy}

		// To make a better grid, we're skipping every other cell along the x axis for roads.
		if building.Kind == constants.BuildingRoad {
			buddy := objects.Renderable{projected.Translate(1, 0), constants.SymbolWallHorizontal, nil}
			buildings = append(buildings, buddy)
		}

		col := color.RGBA{255, 255, 255, 1.0}

		if !building.HasPower() {
			col = color.RGBA{255, 255, 0, 1.0}
		}

		buildings = append(buildings, objects.Renderable{projected, building.Char(), col})
	}

	return buildings
}
