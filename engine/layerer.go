package engine

import "fmt"
import "github.com/dadleyy/charlestown/engine/objects"

type layerer struct {
}

func (l layerer) layer(renderables []objects.Renderable, state objects.Game) []objects.Renderable {
	result := make([]objects.Renderable, 0, len(renderables))
	locations := make(map[string]int)

	for _, item := range renderables {
		key := fmt.Sprintf("%s", &item.Location)
		parent, dupe := locations[key]

		if !dupe {
			parent = 0
		}

		locations[key] = parent + 1

		if !dupe && item.Symbol == ' ' {
			continue
		}

		result = append(result, item)
	}

	return result
}
