package objects

import "fmt"
import "github.com/dadleyy/charlestown/engine/constants"

// Cursor represents the location of the player as well as what mode and building they are currently acting as.
type Cursor struct {
	Location  Point
	Mode      int
	Inventory []Building
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
	case constants.CursorDemolish:
		return constants.SymbolCursorDemolish
	case constants.CursorMove:
		if len(c.Inventory) == 1 {
			b := c.Inventory[0]
			return b.Char()
		}
		return constants.SymbolCursorMove
	default:
		return constants.SymbolCursorNormal
	}
}
