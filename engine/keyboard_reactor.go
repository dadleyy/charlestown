package engine

import "fmt"
import "log"
import "github.com/gdamore/tcell"

// import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/mutations"

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
	quit    chan<- error
	updates chan<- mutations.Mutation
}

// HandleEvent implements the tcell.EventHandler interface for the keyboard reactor.
func (keyboard *keyboardReactor) HandleEvent(event tcell.Event) bool {
	switch event := event.(type) {
	case *tcell.EventKey:
		return keyboard.keyEvent(event)
	default:
		log.Printf("received unknown event, polling next")
	}
	return true
}

func (keyboard *keyboardReactor) runeEvent(key rune) bool {
	switch key {
	case keyUp:
		keyboard.updates <- mutations.Move(0, -verticalMovementSpeed)
	case keyLeft:
		keyboard.updates <- mutations.Move(-horizontalMovementSpeed, 0)
	case keyDown:
		keyboard.updates <- mutations.Move(0, verticalMovementSpeed)
	case keyRight:
		keyboard.updates <- mutations.Move(horizontalMovementSpeed, 0)
	default:
		log.Printf("unknown character key pressed: '%c'", key)
	}
	return true
}

func (keyboard *keyboardReactor) keyEvent(event *tcell.EventKey) bool {
	switch event.Key() {
	case tcell.KeyCtrlC, tcell.KeyEscape:
		log.Printf("exiting by user command")
		keyboard.quit <- fmt.Errorf("user")
		return false
	case tcell.KeyCtrlB:
		keyboard.updates <- mutations.Debug()
	case tcell.KeyBacktab:
		keyboard.updates <- mutations.Intramode()
	case tcell.KeyUp:
		keyboard.updates <- mutations.Move(0, -verticalMovementSpeed)
	case tcell.KeyDown:
		keyboard.updates <- mutations.Move(0, verticalMovementSpeed)
	case tcell.KeyLeft:
		keyboard.updates <- mutations.Move(-horizontalMovementSpeed, 0)
	case tcell.KeyRight:
		keyboard.updates <- mutations.Move(horizontalMovementSpeed, 0)
	case tcell.KeyTab:
		keyboard.updates <- mutations.Mode()
	case tcell.KeyEnter:
		keyboard.updates <- mutations.Interact()
	case tcell.KeyRune:
		return keyboard.runeEvent(event.Rune())
	default:
		log.Printf("unknown keyboard character '%c' (%v): %s", event.Rune(), event.Key(), event.Name())
	}
	return true
}
