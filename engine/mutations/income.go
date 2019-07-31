package mutations

import "time"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

// Income increases the funds currently available to the game based on revenue.
func Income(t time.Time) Mutation {
	return func(state objects.Game) objects.Game {
		next := dup(state)
		dt := time.Now().Sub(t)
		next.Revenue = int(constants.EconomyMultiplier * dt.Seconds())
		next.Funds += next.Revenue
		return next
	}
}
