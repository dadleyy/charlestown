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

type tcellEngine struct {
	config    Configuration
	renderers []renderer
	layerer   layerer
}

func (instance *tcellEngine) draw(screen tcell.Screen, state objects.Game) error {
	screen.Clear()

	if state.Dimensions.Area() < 1 {
		log.Printf("skipping to draw - no dimensions")
		return nil
	}

	renderables := make([]objects.Renderable, 0, state.World.Area())

	for _, renderer := range instance.renderers {
		items := renderer.generate(state)
		renderables = append(renderables, items...)
	}

	layered := instance.layerer.layer(renderables, state)

	for _, r := range layered {
		style := tcell.StyleDefault

		if r.Color != nil {
			red, green, blue, _ := r.Color.RGBA()
			style = tcell.StyleDefault.Foreground(tcell.NewRGBColor(int32(red), int32(green), int32(blue)))
		}

		screen.SetContent(r.Location.X, r.Location.Y, r.Symbol, []rune{}, style)
	}

	screen.Show()
	return nil
}

func (instance *tcellEngine) run(state objects.Game) error {
	log.Printf("[init] initializing encoding")
	encoding.Register()

	log.Printf("[init] creating tcell screen")
	screen, e := tcell.NewScreen()

	if e != nil {
		return e
	}

	log.Printf("[init] initializing screen")
	if e := screen.Init(); e != nil {
		return e
	}

	// Bind some syscall signals to a kill channel.
	log.Printf("[init] registering syscall listeners")
	kills := make(chan os.Signal, 1)
	signal.Notify(kills, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)

	// Create a channel for user input to kill our loop.
	quit := make(chan error)

	// Create the queue of updates we'll receive and apply over.
	updates := make(chan mutations.Mutation, 10)
	wg := &sync.WaitGroup{}

	log.Printf("[init] creating event handlers")
	// Create the list of event handlers that will receive tcell events
	handlers := []tcell.EventHandler{
		&keyboardReactor{quit, updates},
		&viewportReactor{updates},
	}
	multiplex := eventDispatcher{handlers}

	stop := make(chan struct{})
	mutators := []mutator{
		&economyMutator{},
		&flashMutator{},
		&turnMutator{},
	}
	dispatch := timingDispatcher{updates, mutators}

	log.Printf("[init] starting event multiplexer")
	go multiplex.poll(screen, wg)

	log.Printf("[init] starting scheduled mutator dispatch")
	go dispatch.start(stop, wg)

	var exit error
	last := time.Now()

	log.Printf("[init] setting dimensions + first draw")
	width, height := screen.Size()
	state.Dimensions = objects.Dimensions{width, height}
	instance.draw(screen, state)

	log.Printf("[init] entering main game loop")
	for exit == nil {
		select {
		case update := <-updates:
			next := update.Apply(state)

			if len(next.Buildings) != len(state.Buildings) {
				log.Printf("recalculating neighbors")
			}

			state = next
		// terminal states, user
		case <-quit:
			log.Printf("[shutdown] received user shutdown command")
			exit = fmt.Errorf("interrupted")
			break
		case <-kills:
			log.Printf("[shutdown] received shutdown signal, terminating")
			exit = fmt.Errorf("interrupted")
			break
		}

		if exit == nil {
			current := time.Now()
			state.Frame++
			dt := current.Sub(last)

			if s := dt.Seconds(); s >= constants.IdleDelay.Seconds() || state.Frame > 5 {
				instance.draw(screen, state)
				last = current
				state.Frame = 0
			}
		}
	}

	log.Printf("[shutdown] game loop terminated. clearing screen and freeing tcell resources")
	screen.Clear()
	screen.Fini()

	log.Printf("[shutdown] closing economy updater")
	stop <- struct{}{}

	log.Printf("[shutdown] waiting for loop reactors")
	wg.Wait()

	log.Printf("[shutdown] complete")
	return exit
}

func newTcellEngine(config Configuration) engine {
	return &tcellEngine{
		config: config,
		renderers: []renderer{
			&boundaryRenderer{},
			&buildingRenderer{},
			&cursorRenderer{},
			&uiRenderer{},
		},
	}
}
