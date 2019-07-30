package engine

import "log"

type buildingRenderer struct {
	*log.Logger
}

func (renderer *buildingRenderer) generate(state gameState) []renderable {
	buildings := make([]renderable, 0, len(state.buildings))
	hx, hy := state.dimensions.midway()
	cx := state.cursor.location.x
	cy := state.cursor.location.y

	for _, p := range state.buildings {
		projected := point{p.x + hx - cx, p.y + hy - cy}
		renderer.Printf("[house] (original %s) @ %s", p, projected)
		buildings = append(buildings, renderable{projected, symbolHouse})
	}

	return buildings
}
