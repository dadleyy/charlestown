package engine

import "github.com/gdamore/tcell"

type eventSource interface {
	PollEvent() tcell.Event
}
