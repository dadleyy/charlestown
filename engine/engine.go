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

func (instance *engine) draw(screen tcell.Screen, state gameState) error {
	screen.Clear()

	if state.dimensions.height < 1 || state.dimensions.width < 1 {
		instance.Printf("skipping to draw - no dimensions")
		return nil
	}

	instance.Printf("setting cursor %s", state.cursor.location)

	midX := state.dimensions.width / 2
	midY := state.dimensions.height / 2

	screen.SetContent(midX, midY, state.cursor.char(), []rune{}, tcell.StyleDefault)

	left := midX - state.cursor.location.x
	right := left + state.world.width
	top := midY - state.cursor.location.y
	bottom := top + state.world.height

	for i := 1; i < state.world.width; i++ {
		col := i + left
		screen.SetContent(col, top, symbolWallHorizontal, []rune{}, tcell.StyleDefault)
		screen.SetContent(col, bottom, symbolWallHorizontal, []rune{}, tcell.StyleDefault)
	}

	for i := 1; i < state.world.height; i++ {
		row := i + top
		screen.SetContent(left, row, symbolWallVertical, []rune{}, tcell.StyleDefault)
		screen.SetContent(right, row, symbolWallVertical, []rune{}, tcell.StyleDefault)
	}

	screen.SetContent(left, top, symbolWallTopLeft, []rune{}, tcell.StyleDefault)
	screen.SetContent(right, top, symbolWallTopRight, []rune{}, tcell.StyleDefault)
	screen.SetContent(left, bottom, symbolWallBottomLeft, []rune{}, tcell.StyleDefault)
	screen.SetContent(right, bottom, symbolWallBottomRight, []rune{}, tcell.StyleDefault)

	screen.Show()
	return nil
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
	redraw := make(chan mutation)
	wg := &sync.WaitGroup{}
	timer := time.Tick(time.Millisecond * 500)

	instance.Printf("initializing screen")
	if e := screen.Init(); e != nil {
		return e
	}

	instance.Printf("starting keyboard reactor")
	multiplex := eventDispatcher{
		&keyboardReactor{instance.Logger, quit, redraw},
		&viewportReactor{instance.Logger, quit, redraw},
	}

	go multiplex.poll(screen, wg, instance.Logger)

	var exit error

	for exit == nil {
		instance.draw(screen, state)

		select {
		case e := <-quit:
			exit = e
		case update := <-redraw:
			state = update(state)
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
