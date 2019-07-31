package objects

import "fmt"

// Game represents the state of the game.
type Game struct {
	Cursor     Cursor
	Dimensions Dimensions
	World      Dimensions
	Buildings  []Building
	Funds      int
	Revenue    int
	Flashes    []Flash
	Frame      int
}

// String returns a user readable version of the game.
func (state *Game) String() string {
	return fmt.Sprintf("<world %s | window %s | cursor %s>", &state.World, &state.Dimensions, &state.Cursor.Location)
}
