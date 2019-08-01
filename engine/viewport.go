package engine

import "log"
import "github.com/gdamore/tcell"
import "github.com/dadleyy/charlestown/engine/mutations"

type viewportReactor struct {
	*log.Logger
	updates chan<- mutations.Mutation
}

func (viewport *viewportReactor) HandleEvent(event tcell.Event) bool {
	switch event := event.(type) {
	case *tcell.EventResize:
		width, height := event.Size()
		viewport.Printf("resize event, new dimensions (%d, %d)", width, height)
		viewport.updates <- mutations.Resize(width, height)
	}
	return false
}
