package mutations

import "github.com/dadleyy/charlestown/engine/objects"

// Move translates the current game's cursor by the provided x and y amounts.
func Move(x int, y int) Mutation {
	return move{x, y}
}

type move struct {
	x int
	y int
}

func (m move) Apply(game objects.Game) objects.Game {
	next := game
	loc := game.Cursor.Location.Translate(m.x, m.y)

	if loc.X < 1 {
		loc.X = 1
	}

	if loc.Y < 1 {
		loc.Y = 1
	}

	if loc.Y > game.World.Height-1 {
		loc.Y = game.World.Height - 1
	}

	if loc.X > game.World.Width-1 {
		loc.X = game.World.Width - 1
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
