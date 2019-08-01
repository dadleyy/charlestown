package mutations

import "time"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Turn is a mutation that will queue up a new turn.
func Turn(amt int) Mutation {
	return turn{}
}

type turn struct {
}

func (t turn) Apply(game objects.Game) objects.Game {
	homes := 0

	for _, h := range game.Buildings {
		if h.Kind == constants.BuildingHouse {
			homes++
		}
	}

	game.Turn.Actions = objects.TurnActions{0, homes + constants.BaseTurnCardinality}
	game.Turn.Start = time.Now()
	return game
}
