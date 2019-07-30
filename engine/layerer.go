package engine

type layerer struct {
}

func (l layerer) layer(renderables []renderable, state gameState) []renderable {
	result := make([]renderable, 0, len(renderables))

	for _, item := range renderables {
		if item.symbol == ' ' {
			continue
		}

		result = append(result, item)
	}

	return result
}
