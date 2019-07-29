package engine

import "fmt"

const (
	cursorDefault = iota
	cursorBuild
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type cursor struct {
	location point
	kind     int
}

func (c cursor) char() rune {
	switch c.kind {
	case cursorBuild:
		cursorRune = 'O'
	default:
		return 'X'
	}
}

type dimensions struct {
	width  int
	height int
}

type gameState struct {
	cursor     cursor
	dimensions dimensions
}
