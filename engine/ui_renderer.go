package engine

import "log"
import "fmt"
import "time"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/drawing"
import "github.com/dadleyy/charlestown/engine/constants"

type uiRenderer struct {
	*log.Logger
	last time.Time
}

func (renderer *uiRenderer) textWidth(input []string) int {
	max := 0

	for _, t := range input {
		if l := len(t); l > max {
			max = l
		}
	}

	return max
}

func (renderer *uiRenderer) progress(label string, percentage float64, width int) string {
	bar := width - (4 + len(label))
	max := int(float64(bar) * percentage)
	fills := ""

	for i := 0; i < bar; i++ {
		letter := ' '

		if i < max-1 {
			letter = constants.SymbolProgressFill
		}

		fills += fmt.Sprintf("%c", letter)
	}

	return fmt.Sprintf("%s: [%s]", label, fills)
}

func (renderer *uiRenderer) statsMenu(state objects.Game) []objects.Renderable {
	texts := []string{
		fmt.Sprintf("Funds: %d (%+d/s)", state.Funds, state.Revenue),
		fmt.Sprintf("Population: %d (%+d/s)", state.Population, state.PopulationGrowth),
	}

	max := renderer.textWidth(texts)
	actions, timing := state.Turn.Progress()
	texts = append(texts, renderer.progress("Actions", actions, max))
	texts = append(texts, renderer.progress("Time", timing, max))

	return drawing.Box(max, len(texts), texts...)
}

func (renderer *uiRenderer) debugMenu(state objects.Game) []objects.Renderable {
	width, _ := state.Dimensions.Values()
	dt := time.Now().Sub(renderer.last).Seconds()
	fps := 1 / dt
	texts := []string{
		fmt.Sprintf("Debug Information (version %s):", constants.AppVersion),
		fmt.Sprintf("Player location: %s", &state.Cursor.Location),
		fmt.Sprintf("Buildings: %d", len(state.Buildings)),
		fmt.Sprintf("Frame: %d (%.2f f/s)", state.Frame, fps),
		fmt.Sprintf("Flashes: %d", len(state.Flashes)),
		fmt.Sprintf("Mode: %d", state.Cursor.Mode),
	}

	for _, f := range state.Flashes {
		texts = append(texts, fmt.Sprintf("- expiring %.2f", f.Expires.Sub(time.Now()).Seconds()))
	}

	max := renderer.textWidth(texts)
	box := drawing.Box(max, len(texts), texts...)
	return drawing.Translate(box, objects.Point{width - (max + constants.ViewportRightPadding), 0})
}

func (renderer *uiRenderer) flash(flash objects.Flash, pos objects.Point) []objects.Renderable {
	box := drawing.Box(len(flash.Text), 1, flash.Text)
	return drawing.Translate(box, pos)
}

// buildMenu renders out what the user is currently building
func (renderer *uiRenderer) buildMenu(state objects.Game) []objects.Renderable {
	if len(state.Cursor.Inventory) != 1 {
		text := fmt.Sprintf("Build")
		return drawing.Box(len(text), 1, text)
	}

	b := state.Cursor.Inventory[0]
	text := fmt.Sprintf("Build: %c (%s): %d", b.Char(), &b, b.Cost())
	return drawing.Box(len(text), 1, text)
}

func (renderer *uiRenderer) demolishMenu(state objects.Game) []objects.Renderable {
	for _, b := range state.Buildings {
		if b.Location.Equals(state.Cursor.Location) {
			text := fmt.Sprintf("Demolishing: %c (%s)", b.Char(), &b)
			return drawing.Box(len(text), 1, text)
		}
	}

	text := fmt.Sprintf("Demolish")
	return drawing.Box(len(text), 1, text)
}

func (renderer *uiRenderer) moveMenu(state objects.Game) []objects.Renderable {
	carrying := len(state.Cursor.Inventory) == 1

	if carrying {
		b := state.Cursor.Inventory[0]
		text := fmt.Sprintf("Moving: %c (%s)", b.Char(), &b)
		return drawing.Box(len(text), 1, text)
	}

	for _, b := range state.Buildings {
		if b.Location.Equals(state.Cursor.Location) {
			text := fmt.Sprintf("Move: %c (%s)", b.Char(), &b)
			return drawing.Box(len(text), 1, text)
		}
	}

	text := fmt.Sprintf("Move")
	return drawing.Box(len(text), 1, text)
}

func (renderer *uiRenderer) modeMenu(state objects.Game) []objects.Renderable {
	switch state.Cursor.Mode {
	case constants.CursorBuild:
		return renderer.buildMenu(state)
	case constants.CursorMove:
		return renderer.moveMenu(state)
	case constants.CursorDemolish:
		return renderer.demolishMenu(state)
	default:
		return nil
	}
}

func (renderer *uiRenderer) generate(state objects.Game) []objects.Renderable {
	_, height := state.Dimensions.Values()
	result := renderer.statsMenu(state)

	if state.Debug {
		result = append(result, renderer.debugMenu(state)...)
	}

	for index, f := range state.Flashes {
		if index > 2 {
			continue
		}

		transform := objects.Point{0, (height - constants.ViewportBottomPadding) - (index * 3)}
		result = append(result, renderer.flash(f, transform)...)
	}

	renderer.last = time.Now()
	return append(result, drawing.Translate(renderer.modeMenu(state), objects.Point{0, 7})...)
}
