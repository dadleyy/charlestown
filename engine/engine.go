package engine

import "github.com/dadleyy/charlestown/engine/objects"

type engine interface {
	run(objects.Game) error
}
