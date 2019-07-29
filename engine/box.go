package engine

func box(width, height int) []renderable {
	result := make([]renderable, 0, width*height)

	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			symbol := ' '

			if x == width && y == 0 {
				symbol = symbolWallTopRight
			} else if y == 0 && x == 0 {
				symbol = symbolWallTopLeft
			} else if x == 0 && y == height {
				symbol = symbolWallBottomLeft
			} else if x == width && y == height {
				symbol = symbolWallBottomRight
			} else if y == height {
				symbol = symbolWallHorizontal
			} else if x == width {
				symbol = symbolWallVertical
			} else if x == 0 {
				symbol = symbolWallVertical
			} else if y == 0 {
				symbol = symbolWallHorizontal
			}

			result = append(result, renderable{location: point{x, y}, symbol: symbol})
		}
	}

	return result
}
