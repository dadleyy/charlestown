package objects

import "fmt"

// Cursor represents the location of the player as well as what mode and building they are currently acting as.
type Cursor struct {
	Location  Point
	Mode      int
	Inventory []Building
}

// String returns a human redable version of the cursor.
func (c Cursor) String() string {
	return fmt.Sprintf("%s (%d)", c.Location, c.Mode)
}
