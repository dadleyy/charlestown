package engine

import "fmt"

type layerer struct {
}

func (l layerer) layer(renderables []renderable, state gameState) []renderable {
	result := make([]renderable, 0, len(renderables))
	locations := make(map[string]int)

	for _, item := range renderables {
		key := fmt.Sprintf("%s", item.location)
		parent, dupe := locations[key]

		if !dupe {
			parent = 0
		}

		locations[key] = parent + 1

		if !dupe && item.symbol == ' ' {
			continue
		}

		result = append(result, item)
	}

	return result
}
