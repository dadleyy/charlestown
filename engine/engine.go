package engine

import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/mutations"

type engine interface {
	run(objects.Game, <-chan mutations.Mutation) error
}
