package engine

import "log"
import "sync"
import "github.com/gdamore/tcell"

type eventDispatcher struct {
	handlers []tcell.EventHandler
}

func (dispatch *eventDispatcher) poll(screen tcell.Screen, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	event := screen.PollEvent()

	for event != nil {
		for _, d := range dispatch.handlers {
			d.HandleEvent(event)
		}

		event = screen.PollEvent()
	}

	log.Printf("[events] event loop terminated")
}
