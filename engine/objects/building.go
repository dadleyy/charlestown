package objects

import "github.com/dadleyy/charlestown/engine/constants"

// Building instances are like Renderables with additional information.
type Building struct {
	Location Point
	Kind     int
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
	default:
		return "House"
	}
}

// Char returns the symbol to use when rendering the building.
func (b *Building) Char() rune {
	switch b.Kind {
	case constants.BuildingBusiness:
		return constants.SymbolBusiness
	case constants.BuildingRoad:
		return constants.SymbolWallHorizontal
	case constants.BuildingPark:
		return constants.SymbolPark
	default:
		return constants.SymbolHouse
	}
}

// Cost calculates the cost of the building.
func (b *Building) Cost() int {
	switch b.Kind {
	case constants.BuildingBusiness:
		return 10 * constants.EconomyMultiplier
	case constants.BuildingRoad:
		return 2 * constants.EconomyMultiplier
	case constants.BuildingPark:
		return 2 * constants.EconomyMultiplier
	default:
		return 4 * constants.EconomyMultiplier
	}
}
