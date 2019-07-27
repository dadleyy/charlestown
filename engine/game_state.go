package engine

const (
	cursorDefault = iota
	cursorBuild
)

type cursor struct {
	x    int
	y    int
	kind int
}

type dimensions struct {
	width  int
	height int
}

type gameState struct {
	cursor     cursor
	dimensions dimensions
}
