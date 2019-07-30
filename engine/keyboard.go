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

	horizontalMovementSpeed = 2
	verticalMovementSpeed   = 1
)

type keyboardReactor struct {
	*log.Logger
	quit    chan<- error
	updates chan<- mutation
}

// HandleEvent implements the tcell.EventHandler interface for the keyboard reactor.
func (keyboard *keyboardReactor) HandleEvent(event tcell.Event) bool {
	switch event := event.(type) {
	case *tcell.EventKey:
		return keyboard.keyEvent(event)
	default:
		keyboard.Printf("received unknown event, polling next")
	}
	return true
}

func (keyboard *keyboardReactor) runeEvent(key rune) bool {
	switch key {
	case keyUp:
		keyboard.updates <- move(0, -verticalMovementSpeed)
	case keyLeft:
		keyboard.updates <- move(-horizontalMovementSpeed, 0)
	case keyDown:
		keyboard.updates <- move(0, verticalMovementSpeed)
	case keyRight:
		keyboard.updates <- move(horizontalMovementSpeed, 0)
	default:
		keyboard.Printf("unknown character key pressed: '%c'", key)
	}
	return true
}

func (keyboard *keyboardReactor) keyEvent(event *tcell.EventKey) bool {
	switch event.Key() {
	case tcell.KeyCtrlC, tcell.KeyEscape:
		keyboard.Printf("exiting by user command")
		keyboard.quit <- fmt.Errorf("user")
		return false
	case tcell.KeyBacktab:
		keyboard.updates <- intramode()
	case tcell.KeyUp:
		keyboard.updates <- move(0, -verticalMovementSpeed)
	case tcell.KeyDown:
		keyboard.updates <- move(0, verticalMovementSpeed)
	case tcell.KeyLeft:
		keyboard.updates <- move(-horizontalMovementSpeed, 0)
	case tcell.KeyRight:
		keyboard.updates <- move(horizontalMovementSpeed, 0)
	case tcell.KeyTab:
		keyboard.updates <- mode()
	case tcell.KeyEnter:
		keyboard.updates <- interact()
	case tcell.KeyRune:
		return keyboard.runeEvent(event.Rune())
	default:
		keyboard.Printf("unknown keyboard character '%c' (%v): %s", event.Rune(), event.Key(), event.Name())
	}
	return true
}
