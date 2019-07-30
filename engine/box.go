package engine

import "unicode/utf8"

func box(width, height int, text ...string) []renderable {
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

			if len(text) > y-1 && y > 0 {
				message := text[y-1]

				if x > 0 {
					r, s := utf8.DecodeRuneInString(message)

					if s > 0 {
						symbol = r
						text[y-1] = message[s:]
					}
				}
			}

			result = append(result, renderable{location: point{x, y}, symbol: symbol})
		}
	}

	return result
}
