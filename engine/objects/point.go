package objects

import "fmt"

// Point wraps an x and y position.
type Point struct {
	X int
	Y int
}

// String returns a human readable string representation of the point.
func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Dup copies the point.
func (p *Point) Dup() Point {
	return Point{X: p.X, Y: p.Y}
}

// Translate returns a new point having the x and y coordinates moved by the provided amounts.
func (p *Point) Translate(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}
