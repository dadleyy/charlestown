package engine

const (
	buildingHouse = iota
	buildingPark
	buildingBusiness
)

type building struct {
	location point
	kind     int
}

func (b *building) char() rune {
	switch b.kind {
	case buildingBusiness:
		return symbolBusiness
	case buildingPark:
		return symbolPark
	default:
		return symbolHouse
	}
}
