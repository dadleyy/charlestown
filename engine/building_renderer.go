package engine

import "log"

const (
	neighborWest  = "west"
	neighborNorth = "north"
	neighborEast  = "east"
	neighborSouth = "south"
)

type buildingRenderer struct {
	*log.Logger
}

func (renderer *buildingRenderer) neighbors(p point, all []building) map[string]building {
	result := make(map[string]building)

	for _, b := range all {
		if b.location.x == p.x-2 && b.location.y == p.y {
			result[neighborWest] = b
		}
		if b.location.x == p.x && b.location.y == p.y-1 {
			result[neighborNorth] = b
		}
		if b.location.x == p.x && b.location.y == p.y+1 {
			result[neighborSouth] = b
		}
		if b.location.x == p.x+2 && b.location.y == p.y {
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

func (renderer *buildingRenderer) symbolForRoad(neighbors map[string]building) rune {
	_, n := neighbors[neighborNorth]
	_, e := neighbors[neighborEast]
	_, s := neighbors[neighborSouth]
	_, w := neighbors[neighborWest]
	size := len(neighbors)

	if renderer.isCross(size, n, e, s, w) {
		return symbolWallCross
	} else if renderer.isBottomMid(size, n, e, s, w) {
		return symbolWallWestSouthEast
	} else if renderer.isBottomLeft(size, n, e, s, w) {
		return symbolWallBottomLeft
	} else if renderer.isTopMid(size, n, e, s, w) {
		return symbolWallWestNorthEast
	} else if (n && s) || (n && size == 1) || (s && size == 1) {
		return symbolWallVertical
	}

	return symbolWallHorizontal
}

func (renderer *buildingRenderer) generate(state gameState) []renderable {
	buildings := make([]renderable, 0, len(state.buildings))
	hx, hy := state.dimensions.midway()
	cx := state.cursor.location.x
	cy := state.cursor.location.y

	for _, p := range state.buildings {
		projected := point{p.location.x + hx - cx, p.location.y + hy - cy}

		if p.kind == buildingRoad {
			neighbors := renderer.neighbors(p.location, state.buildings)
			buddy := renderable{point{projected.x + 1, projected.y}, p.char()}
			self := renderable{projected, renderer.symbolForRoad(neighbors)}
			buildings = append(buildings, buddy, self)
			continue
		}

		buildings = append(buildings, renderable{projected, p.char()})
	}

	return buildings
}
