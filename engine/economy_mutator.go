package engine

import "time"
import "github.com/dadleyy/charlestown/engine/mutations"

type economyMutator struct {
	delay time.Duration
}

func (mutator *economyMutator) tick(t time.Duration) []mutations.Mutation {
	mutator.delay += t

	if mutator.delay < time.Second {
		return []mutations.Mutation{}
	}

	mutator.delay = 0
	return []mutations.Mutation{mutations.Income(time.Now().Add(-t))}
}
