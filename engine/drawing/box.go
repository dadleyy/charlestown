package drawing

import "unicode/utf8"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

func boxSymbol(x, y, width, height int) rune {
	symbol := ' '

	if x == width && y == 0 {
		symbol = constants.SymbolWallTopRight
	} else if y == 0 && x == 0 {
		symbol = constants.SymbolWallTopLeft
	} else if x == 0 && y == height {
		symbol = constants.SymbolWallBottomLeft
	} else if x == width && y == height {
		symbol = constants.SymbolWallBottomRight
	} else if y == height {
		symbol = constants.SymbolWallHorizontal
	} else if x == width {
		symbol = constants.SymbolWallVertical
	} else if x == 0 {
		symbol = constants.SymbolWallVertical
	} else if y == 0 {
		symbol = constants.SymbolWallHorizontal
	}

	return symbol
}

// Box creates an array of renderables of the provided dimensions with optional text rows.
func Box(width, height int, text ...string) []objects.Renderable {
	result := make([]objects.Renderable, 0, width*height)

	// nolint: gocyclo
	for x := 0; x <= width+1; x++ {
		for y := 0; y <= height+1; y++ {
			symbol := boxSymbol(x, y, width+1, height+1)

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

			result = append(result, objects.Renderable{Location: objects.Point{x, y}, Symbol: symbol})
		}
	}

	return result
}
