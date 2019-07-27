package engine

import "fmt"
import "log"
import "sync"
import "github.com/gdamore/tcell"

type keyboardReactor struct {
	*log.Logger
	quit    chan<- error
	updates chan<- mutation
	wait    *sync.WaitGroup
}

func (keyboard *keyboardReactor) poll(source eventSource) {
	keyboard.wait.Add(1)
	defer keyboard.wait.Done()
	keyboard.Printf("starting keyboard reactor loop")

	event := source.PollEvent()

	for event != nil {
		switch event := event.(type) {
		case *tcell.EventKey:
			keyboard.Printf("received keyboard event")

			switch event.Key() {
			case tcell.KeyCtrlC, tcell.KeyEscape:
				keyboard.Printf("exiting by user command")
				keyboard.quit <- fmt.Errorf("user")
				return
			case tcell.KeyRune:
				switch event.Rune() {
				case 'b':
					keyboard.Printf("toggling build")
					keyboard.updates <- cursorChange(cursorBuild)
				case 'w':
					keyboard.Printf("moving up")
					keyboard.updates <- move(0, -1)
				case 'a':
					keyboard.Printf("moving lest")
					keyboard.updates <- move(-1, 0)
				case 's':
					keyboard.Printf("moving down")
					keyboard.updates <- move(0, 1)
				case 'd':
					keyboard.Printf("moving right")
					keyboard.updates <- move(1, 0)
				default:
					keyboard.Printf("character key pressed: '%c'", event.Rune())
				}
			default:
				keyboard.Printf("unknown keyboard character '%c' (%v)", event.Rune(), event.Key())
			}
		case *tcell.EventResize:
			width, height := event.Size()
			keyboard.Printf("resize event, new dimensions (%d, %d)", width, height)
			keyboard.updates <- resize(width, height)
		default:
			keyboard.Printf("received unknown event, polling next")
		}

		event = source.PollEvent()
	}

	keyboard.Printf("keyboard reactor poll complete")
}
