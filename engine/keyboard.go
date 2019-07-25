package engine

import "fmt"
import "log"
import "sync"
import "github.com/gdamore/tcell"

type keyboardReactor struct {
	*log.Logger
	quit    chan<- error
	updates chan<- gameState
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
			default:
				keyboard.Printf("unknown keyboard character '%c' (%v)", event.Rune(), event.Key())
			}

		default:
			keyboard.Printf("received unknown event, polling next")
		}

		event = source.PollEvent()
	}

	keyboard.Printf("keyboard reactor poll complete")
}
