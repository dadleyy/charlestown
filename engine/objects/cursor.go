package objects

import "fmt"
import "github.com/dadleyy/charlestown/engine/constants"

// Cursor represents the location of the player as well as what mode and building they are currently acting as.
type Cursor struct {
	Location Point
	Mode     int
	Building int
}

// String returns a human redable version of the cursor.
func (c *Cursor) String() string {
	return fmt.Sprintf("%s", &c.Location)
}

// Char returns the rune to use when rendering the cursor.
func (c *Cursor) Char() rune {
	switch c.Mode {
	case constants.CursorBuild:
		return constants.SymbolCursorBuild
	default:
		return constants.SymbolCursorNormal
	}
}

// Dup copies the current cursor, returning a fresh instance.
func (c *Cursor) Dup() Cursor {
	return Cursor{
		Location: c.Location.Dup(),
		Mode:     c.Mode,
		Building: c.Building,
	}
}
