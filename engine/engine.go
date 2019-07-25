package engine

import "os"
import "fmt"
import "log"
import "time"
import "sync"
import "syscall"
import "os/signal"
import "github.com/gdamore/tcell"
import "github.com/gdamore/tcell/encoding"

type engine struct {
	*log.Logger
	config Configuration
}

func (instance *engine) run(state gameState) error {
	instance.Printf("initializing encoding")
	encoding.Register()
	screen, e := tcell.NewScreen()

	if e != nil {
		return e
	}

	kills := make(chan os.Signal, 1)
	signal.Notify(kills, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)

	quit := make(chan error)
	redraw := make(chan gameState)
	wg := &sync.WaitGroup{}
	timer := time.Tick(time.Millisecond * 500)

	instance.Printf("initializing screen")
	if e := screen.Init(); e != nil {
		return e
	}

	instance.Printf("starting keyboard reactor")
	keyboard := keyboardReactor{
		Logger:  instance.Logger,
		quit:    quit,
		updates: redraw,
		wait:    wg,
	}

	go keyboard.poll(screen)

	var exit error

	for exit == nil {
		select {
		case e := <-quit:
			exit = e
		case update := <-redraw:
			state = update
			instance.Printf("redrawing game with state %v", state)
		case <-timer:
			instance.Printf("applying update")
		case <-kills:
			instance.Printf("received shutdown signal, terminating")
			exit = fmt.Errorf("interrupted")
		}
	}

	instance.Printf("game loop terminated, clearing screen and closing buffers")
	screen.Clear()
	screen.Fini()
	instance.Printf("waiting for loop reactors")
	wg.Wait()
	instance.Printf("terminating")
	return exit
}
