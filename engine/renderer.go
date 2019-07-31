package engine

import "github.com/dadleyy/charlestown/engine/objects"

type renderer interface {
	generate(objects.Game) []objects.Renderable
}
