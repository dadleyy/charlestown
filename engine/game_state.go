package engine

import "fmt"

type gameState struct {
	cursor     cursor
	dimensions dimensions
	world      dimensions
	buildings  []building
	funds      int
}

func (state *gameState) String() string {
	return fmt.Sprintf("<world %s | window %s | cursor %s>", &state.world, &state.dimensions, &state.cursor.location)
}
