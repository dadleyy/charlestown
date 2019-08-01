package objects

import "fmt"

// Renderable represents a location and a symbol to be rendered.
type Renderable struct {
	Location Point
	Symbol   rune
}

// String returns a human readable version of the renderable.
func (r *Renderable) String() string {
	return fmt.Sprintf("<'%s'@%c>", &r.Location, r.Symbol)
}
