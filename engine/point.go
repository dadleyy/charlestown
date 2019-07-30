package engine

import "fmt"

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}
