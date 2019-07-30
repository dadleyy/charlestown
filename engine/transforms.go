package engine

func translate(items []renderable, offset point) []renderable {
	out := make([]renderable, 0, len(items))

	for _, item := range items {
		next := point{item.location.x + offset.x, item.location.y + offset.y}
		out = append(out, renderable{symbol: item.symbol, location: next})
	}

	return out
}
