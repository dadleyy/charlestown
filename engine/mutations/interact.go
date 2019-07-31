package mutations

import "log"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Interact mutates the state based on the current mode.
func Interact(l *log.Logger) Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)

		if next.Cursor.Mode == constants.CursorBuild {
			l.Printf("adding new building at %s", &next.Cursor.Location)
			construct := objects.Building{
				Location: state.Cursor.Location,
				Kind:     state.Cursor.Building,
			}

			if construct.Cost() > next.Funds {
				return next
			}

			next.Funds -= construct.Cost()
			next.Buildings = append(next.Buildings, construct)
		}

		return next
	}
}
