package drawing

import "github.com/dadleyy/charlestown/engine/objects"

// Translate maps each provided renderable to a new location based on a provided offset.
func Translate(items []objects.Renderable, offset objects.Point) []objects.Renderable {
	out := make([]objects.Renderable, 0, len(items))

	for _, item := range items {
		next := item.Location.Translate(offset.X, offset.Y)
		out = append(out, objects.Renderable{Symbol: item.Symbol, Location: next})
	}

	return out
}
