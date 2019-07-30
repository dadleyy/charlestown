package engine

const (
	cursorNormal = iota
	cursorBuild
)

type cursor struct {
	location point
	mode     int
	building int
}

func (c cursor) char() rune {
	switch c.mode {
	case cursorBuild:
		return symbolCursorBuild
	default:
		return symbolCursorNormal
	}
}
