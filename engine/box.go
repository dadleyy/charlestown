package engine

import "unicode/utf8"

func boxSymbol(x, y, width, height int) rune {
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

	return symbol
}

func box(width, height int, text ...string) []renderable {
	result := make([]renderable, 0, width*height)

	// nolint: gocyclo
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			symbol := boxSymbol(x, y, width, height)

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
