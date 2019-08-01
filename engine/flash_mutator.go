package engine

import "time"
import "github.com/dadleyy/charlestown/engine/mutations"

type flashMutator struct {
	delay time.Duration
}

func (mutator *flashMutator) tick(d time.Duration) []mutations.Mutation {
	mutator.delay += d

	if mutator.delay < time.Millisecond*500 {
		return nil
	}

	mutator.delay = 0
	return []mutations.Mutation{mutations.Expire()}
}
