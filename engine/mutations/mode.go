package mutations

import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Mode toggles the game cursor's current mode.
func Mode() Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)
		switch state.Cursor.Mode {
		case constants.CursorNormal:
			next.Cursor.Mode = constants.CursorBuild
		default:
			next.Cursor.Mode = constants.CursorNormal
			next.Cursor.Building = constants.BuildingHouse
		}
		return next
	}
}

// Intramode interacts with the current state based on the mode.
func Intramode() Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)

		if next.Cursor.Mode == constants.CursorBuild {
			target := next.Cursor.Building + 1
			next.Cursor.Building = target % (constants.BuildingBusiness + 1)
		}

		return next
	}
}
