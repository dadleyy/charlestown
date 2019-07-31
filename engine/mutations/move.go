package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Move translates the current game's cursor by the provided x and y amounts.
func Move(x int, y int) Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)
		loc := state.Cursor.Location.Translate(x, y)

		if loc.X < 1 {
			loc.X = 1
		}

		if loc.Y < 1 {
			loc.Y = 1
		}

		if loc.Y > state.World.Height-1 {
			loc.Y = state.World.Height - 1
		}

		if loc.X > state.World.Width-1 {
			loc.X = state.World.Width - 1
		}

		if loc.Y < 0 {
			loc.Y = 0
		}

		if loc.X < 0 {
			loc.X = 0
		}

		next.Cursor.Location = loc
		return next
	}
}
