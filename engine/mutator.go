package engine

import "time"
import "github.com/dadleyy/charlestown/engine/mutations"

type mutator interface {
	tick(time.Duration) []mutations.Mutation
}
