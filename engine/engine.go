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
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/mutations"
import "github.com/dadleyy/charlestown/engine/constants"

type engine struct {
	*log.Logger
	config    Configuration
	renderers []renderer
	layerer   layerer
}

func (instance *engine) draw(screen tcell.Screen, state objects.Game) error {
	screen.Clear()

	if state.Dimensions.Area() < 1 {
		instance.Printf("skipping to draw - no dimensions")
		return nil
	}

	renderables := make([]objects.Renderable, 0, state.World.Area())

	for _, renderer := range instance.renderers {
		items := renderer.generate(state)
		renderables = append(renderables, items...)
	}

	layered := instance.layerer.layer(renderables, state)

	for _, r := range layered {
		screen.SetContent(r.Location.X, r.Location.Y, r.Symbol, []rune{}, tcell.StyleDefault)
	}

	screen.Show()
	return nil
}

func (instance *engine) run(state objects.Game) error {
	instance.Printf("[init] initializing encoding")
	encoding.Register()

	instance.Printf("[init] creating tcell screen")
	screen, e := tcell.NewScreen()

	if e != nil {
		return e
	}

	instance.Printf("[init] initializing screen")
	if e := screen.Init(); e != nil {
		return e
	}

	// Bind some syscall signals to a kill channel.
	instance.Printf("[init] registering syscall listeners")
	kills := make(chan os.Signal, 1)
	signal.Notify(kills, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)

	// Create a channel for user input to kill our loop.
	quit := make(chan error)

	// Create the queue of updates we'll receive and apply over.
	updates := make(chan mutations.Mutation, 10)
	wg := &sync.WaitGroup{}

	instance.Printf("[init] creating event handlers")
	// Create the list of event handlers that will receive tcell events
	handlers := []tcell.EventHandler{
		&keyboardReactor{instance.Logger, quit, updates},
		&viewportReactor{instance.Logger, updates},
	}
	multiplex := eventDispatcher{instance.Logger, handlers}

	stop := make(chan struct{})
	econ := economyManager{instance.Logger, updates}

	instance.Printf("[init] starting event multiplexer")
	go multiplex.poll(screen, wg)

	instance.Printf("[init] starting economy manager")
	go econ.tick(stop, wg)

	var exit error
	last := time.Now()

	instance.Printf("[init] setting dimensions + first draw")
	width, height := screen.Size()
	state.Dimensions = objects.Dimensions{width, height}
	instance.draw(screen, state)

	instance.Printf("[init] entering main game loop")
	for exit == nil {
		state.Frame++

		select {
		case update := <-updates:
			state = update(state)
		// terminal states, user
		case <-quit:
			instance.Printf("[shutdown] received user shutdown command")
			exit = fmt.Errorf("interrupted")
			break
		case <-kills:
			instance.Printf("[shutdown] received shutdown signal, terminating")
			exit = fmt.Errorf("interrupted")
			break
		}

		if exit == nil {
			current := time.Now()
			dt := current.Sub(last)

			if s := dt.Seconds(); s >= constants.IdleDelay.Seconds() {
				instance.draw(screen, state)
				last = current
			}
		}
	}

	instance.Printf("[shutdown] game loop terminated. clearing screen and freeing tcell resources")
	screen.Clear()
	screen.Fini()

	instance.Printf("[shutdown] closing economy updater")
	stop <- struct{}{}

	instance.Printf("[shutdown] waiting for loop reactors")
	wg.Wait()

	instance.Printf("[shutdown] complete")
	return exit
}
