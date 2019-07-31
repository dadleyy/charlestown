package engine

import "log"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

const (
	neighborWest  = "west"
	neighborNorth = "north"
	neighborEast  = "east"
	neighborSouth = "south"
)

type buildingRenderer struct {
	*log.Logger
}

func (renderer *buildingRenderer) neighbors(p objects.Point, all []objects.Building) map[string]objects.Building {
	result := make(map[string]objects.Building)

	for _, b := range all {
		if b.Location.X == p.X-2 && b.Location.Y == p.Y {
			result[neighborWest] = b
		}
		if b.Location.X == p.X && b.Location.Y == p.Y-1 {
			result[neighborNorth] = b
		}
		if b.Location.X == p.X && b.Location.Y == p.Y+1 {
			result[neighborSouth] = b
		}
		if b.Location.X == p.X+2 && b.Location.Y == p.Y {
			result[neighborEast] = b
		}
	}

	return result
}

func (renderer *buildingRenderer) isCross(size int, n, e, s, w bool) bool {
	return size == 4 || (size == 2 && n && w) || (size == 3 && n && s && w)
}

func (renderer *buildingRenderer) isBottomMid(size int, n, e, s, w bool) bool {
	return (size == 2 && s && w) || (size == 2 && s && e) || (size == 3 && s && e && w)
}

func (renderer *buildingRenderer) isBottomLeft(size int, n, e, s, w bool) bool {
	return size == 2 && n && e
}

func (renderer *buildingRenderer) isTopMid(size int, n, e, s, w bool) bool {
	return ((size == 2) && n && w) || (size == 3) && n && e && w
}

func (renderer *buildingRenderer) symbolForRoad(neighbors map[string]objects.Building) rune {
	_, n := neighbors[neighborNorth]
	_, e := neighbors[neighborEast]
	_, s := neighbors[neighborSouth]
	_, w := neighbors[neighborWest]
	size := len(neighbors)

	if renderer.isCross(size, n, e, s, w) {
		return constants.SymbolWallCross
	} else if renderer.isBottomMid(size, n, e, s, w) {
		return constants.SymbolWallWestSouthEast
	} else if renderer.isBottomLeft(size, n, e, s, w) {
		return constants.SymbolWallBottomLeft
	} else if renderer.isTopMid(size, n, e, s, w) {
		return constants.SymbolWallWestNorthEast
	} else if (n && s) || (n && size == 1) || (s && size == 1) {
		return constants.SymbolWallVertical
	}

	return constants.SymbolWallHorizontal
}

func (renderer *buildingRenderer) generate(state objects.Game) []objects.Renderable {
	buildings := make([]objects.Renderable, 0, len(state.Buildings))

	hx, hy := state.Dimensions.Midway()
	cx := state.Cursor.Location.X
	cy := state.Cursor.Location.Y

	for _, p := range state.Buildings {
		projected := objects.Point{p.Location.X + hx - cx, p.Location.Y + hy - cy}

		if p.Kind == constants.BuildingRoad {
			// Find our neighbors
			neighbors := renderer.neighbors(p.Location, state.Buildings)

			// To make a better grid, we're skipping every other cell along the x axis
			buddy := objects.Renderable{objects.Point{projected.X + 1, projected.Y}, p.Char()}

			self := objects.Renderable{projected, renderer.symbolForRoad(neighbors)}
			buildings = append(buildings, buddy, self)
			continue
		}

		buildings = append(buildings, objects.Renderable{projected, p.Char()})
	}

	return buildings
}
