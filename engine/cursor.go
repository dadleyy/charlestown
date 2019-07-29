package engine

const (
	cursorDefault = iota
	cursorBuild
)

type cursor struct {
	location point
	kind     int
}

func (c cursor) char() rune {
	switch c.kind {
	case cursorBuild:
		return symbolCursorBuild
	default:
		return symbolCursorNormal
	}
}
