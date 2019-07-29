package engine

import "log"
import "sync"
import "github.com/gdamore/tcell"

type eventDispatcher []tcell.EventHandler

func (dispatch *eventDispatcher) poll(screen tcell.Screen, wg *sync.WaitGroup, log *log.Logger) {
	wg.Add(1)
	defer wg.Done()
	event := screen.PollEvent()

	for event != nil {
		log.Printf("received event %v", event)

		for _, d := range *dispatch {
			d.HandleEvent(event)
		}

		event = screen.PollEvent()
	}
}
