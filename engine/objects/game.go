package objects

import "fmt"

// Game represents the state of the game.
type Game struct {
	Cursor           Cursor
	Dimensions       Dimensions
	World            Dimensions
	Buildings        []Building
	Turn             Turn
	Funds            int
	Revenue          int
	Population       int
	PopulationGrowth int
	Flashes          []Flash
	Frame            int
	Debug            bool
}

// String returns a user readable version of the game.
func (state Game) String() string {
	return fmt.Sprintf(
		"<world %s | window %s | cursor %s>",
		state.World,
		state.Dimensions,
		state.Cursor,
	)
}
