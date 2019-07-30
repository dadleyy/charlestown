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
	config    Configuration
	renderers []renderer
	layerer   layerer
}

func (instance *engine) draw(screen tcell.Screen, state gameState) error {
	screen.Clear()

	if state.dimensions.height < 1 || state.dimensions.width < 1 {
		instance.Printf("skipping to draw - no dimensions")
		return nil
	}

	renderables := make([]renderable, 0, state.world.width*state.world.height)

	for _, renderer := range instance.renderers {
		items := renderer.generate(state)
		renderables = append(renderables, items...)
	}

	layered := instance.layerer.layer(renderables, state)
	instance.Printf("rendering %d points", len(layered))

	for _, r := range layered {
		screen.SetContent(r.location.x, r.location.y, r.symbol, []rune{}, tcell.StyleDefault)
	}

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
