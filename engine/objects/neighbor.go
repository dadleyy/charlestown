package objects

// Neighbor is used to create an array of buildings that are adjacent to a given building.
type Neighbor struct {
	Building  Building
	Direction string
}
