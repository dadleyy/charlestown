package engine

import "log"
import "github.com/gdamore/tcell"

type viewportReactor struct {
	*log.Logger
	quit    chan<- error
	updates chan<- mutation
}

func (viewport *viewportReactor) HandleEvent(event tcell.Event) bool {
	switch event := event.(type) {
	case *tcell.EventResize:
		width, height := event.Size()
		viewport.Printf("resize event, new dimensions (%d, %d)", width, height)
		viewport.updates <- resize(width, height)
	}
	return false
}
