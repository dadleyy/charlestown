package engine

import "time"
import "github.com/dadleyy/charlestown/engine/constants"
import "github.com/dadleyy/charlestown/engine/mutations"

type turnMutator struct {
	delay time.Duration
}

func (t *turnMutator) tick(dt time.Duration) []mutations.Mutation {
	t.delay += dt

	if t.delay < constants.TurnDuration {
		return []mutations.Mutation{}
	}

	t.delay = 0
	return []mutations.Mutation{mutations.Turn(1)}
}
