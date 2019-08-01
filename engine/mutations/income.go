package mutations

import "time"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Income increases the funds currently available to the game based on revenue.
func Income(t time.Time) Mutation {
	return income{t}
}

type income struct {
	created time.Time
}

func (i income) Apply(game objects.Game) objects.Game {
	next := game
	dt := time.Now().Sub(i.created)
	revenue := 0

	for _, b := range next.Buildings {
		revenue += b.Revenue()
	}

	next.Revenue = revenue + int(constants.EconomyMultiplier*dt.Seconds())
	next.Funds += next.Revenue
	return next
}
