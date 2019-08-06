package objects

import "fmt"

// Point wraps an x and y position.
type Point struct {
	X int
	Y int
}

// String returns a human readable string representation of the point.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Translate returns a new point having the x and y coordinates moved by the provided amounts.
func (p Point) Translate(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

// Equals returns true if the coordinates match
func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

// Values returns the x and y coordinates as a tuple
func (p Point) Values() (int, int) {
	return p.X, p.Y
}
