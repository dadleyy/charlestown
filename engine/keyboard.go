package engine

import "fmt"
import "log"
import "github.com/gdamore/tcell"

const (
	keyToggleBuild = 'b'
	keyUp          = 'w'
	keyDown        = 's'
	keyLeft        = 'a'
	keyRight       = 'd'
)

type keyboardReactor struct {
	*log.Logger
	quit    chan<- error
	updates chan<- mutation
}

func (keyboard *keyboardReactor) HandleEvent(event tcell.Event) bool {
	switch event := event.(type) {
	case *tcell.EventKey:
		keyboard.Printf("received keyboard event")

		switch event.Key() {
		case tcell.KeyCtrlC, tcell.KeyEscape:
			keyboard.Printf("exiting by user command")
			keyboard.quit <- fmt.Errorf("user")
			return false
		case tcell.KeyUp:
			keyboard.updates <- move(0, -1)
		case tcell.KeyDown:
			keyboard.updates <- move(0, 1)
		case tcell.KeyLeft:
			keyboard.updates <- move(-1, 0)
		case tcell.KeyRight:
			keyboard.updates <- move(1, 0)
		case tcell.KeyTab:
			keyboard.updates <- mode()
		case tcell.KeyRune:
			switch event.Rune() {
			case keyUp:
				keyboard.updates <- move(0, -1)
			case keyLeft:
				keyboard.updates <- move(-1, 0)
			case keyDown:
				keyboard.updates <- move(0, 1)
			case keyRight:
				keyboard.updates <- move(1, 0)
			default:
				keyboard.Printf("character key pressed: '%c'", event.Rune())
			}
		default:
			keyboard.Printf("unknown keyboard character '%c' (%v)", event.Rune(), event.Key())
		}
	default:
		keyboard.Printf("received unknown event, polling next")
	}
	return true
}
