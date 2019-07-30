package engine

const (
	buildingHouse = iota
	buildingPark
	buildingRoad
	buildingBusiness
)

type building struct {
	location point
	kind     int
}

func (b *building) String() string {
	switch b.kind {
	case buildingBusiness:
		return "Business"
	case buildingPark:
		return "Park"
	case buildingRoad:
		return "Road"
	default:
		return "House"
	}
}

func (b *building) char() rune {
	switch b.kind {
	case buildingBusiness:
		return symbolBusiness
	case buildingRoad:
		return symbolWallHorizontal
	case buildingPark:
		return symbolPark
	default:
		return symbolHouse
	}
}

func (b *building) cost() int {
	switch b.kind {
	case buildingBusiness:
		return 10
	case buildingRoad:
		return 2
	case buildingPark:
		return 2
	default:
		return 4
	}
}
