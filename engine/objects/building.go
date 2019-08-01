package objects

import "github.com/dadleyy/charlestown/engine/constants"

// Building instances are like Renderables with additional information.
type Building struct {
	Location  Point
	Kind      int
	Neighbors []Neighbor
}

// String returns the humanized name of the building.
func (b *Building) String() string {
	switch b.Kind {
	case constants.BuildingBusiness:
		return "Business"
	case constants.BuildingPark:
		return "Park"
	case constants.BuildingRoad:
		return "Road"
	case constants.BuildingPower:
		return "Power"
	default:
		return "House"
	}
}

func (b *Building) neighborPresence() (bool, bool, bool, bool, bool) {
	mapped := make(map[string]Building)

	for _, n := range b.Neighbors {
		mapped[n.Direction] = n.Building

		if n.Building.Kind == constants.BuildingPower {
			mapped[constants.NeighborPower] = n.Building
		}
	}

	_, n := mapped[constants.NeighborNorth]
	_, e := mapped[constants.NeighborEast]
	_, s := mapped[constants.NeighborSouth]
	_, w := mapped[constants.NeighborWest]
	_, p := mapped[constants.NeighborPower]

	return n, e, s, w, p
}

func (b *Building) roadChar() rune {
	n, e, s, w, _ := b.neighborPresence()

	if len(b.Neighbors) == 4 {
		return constants.SymbolWallCross
	}

	if n && e && s {
		return constants.SymbolWallNorthEastSouth
	} else if n && w && s {
		return constants.SymbolWallCross
	} else if w && s && e {
		return constants.SymbolWallWestSouthEast
	} else if n && e {
		return constants.SymbolWallWestNorthEast
	} else if n && w {
		return constants.SymbolWallWestNorthEast
	} else if n && s {
		return constants.SymbolWallNorthEastSouth
	} else if s && w {
		return constants.SymbolWallWestSouthEast
	} else if n {
		return constants.SymbolWallBottomLeft
	} else if s {
		return constants.SymbolWallTopLeft
	}

	return constants.SymbolWallHorizontal
}

// Char returns the symbol to use when rendering the building.
func (b *Building) Char() rune {
	switch b.Kind {
	case constants.BuildingBusiness:
		return constants.SymbolBusiness
	case constants.BuildingRoad:
		return b.roadChar()
	case constants.BuildingPark:
		return constants.SymbolPark
	case constants.BuildingPower:
		return constants.SymbolPower
	default:
		return constants.SymbolHouse
	}
}

// Revenue calculates the revenue of the building (assuming power)
func (b *Building) Revenue() int {
	powerModifier := 0

	if b.HasPower() {
		powerModifier = 1
	}

	switch b.Kind {
	case constants.BuildingBusiness:
		return constants.BuildingRevenueBusiness * powerModifier
	case constants.BuildingPower:
		return constants.BuildingRevenuePower
	case constants.BuildingRoad:
		return constants.BuildingRevenueRoad
	case constants.BuildingPark:
		return constants.BuildingRevenuePark
	default:
		return constants.BuildingRevenueHouse * powerModifier
	}
}

// Cost calculates the cost of the building.
func (b *Building) Cost() int {
	switch b.Kind {
	case constants.BuildingBusiness:
		return constants.BuildingCostBusiness
	case constants.BuildingPower:
		return constants.BuildingCostPower
	case constants.BuildingRoad:
		return constants.BuildingCostRoad
	case constants.BuildingPark:
		return constants.BuildingCostPark
	default:
		return constants.BuildingCostHouse
	}
}

// HasPower will return true if any neighbor has power.
func (b *Building) HasPower() bool {
	switch b.Kind {
	case constants.BuildingRoad, constants.BuildingPower:
		return true
	default:
		_, _, _, _, p := b.neighborPresence()
		return p
	}
}
