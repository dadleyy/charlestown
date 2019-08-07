package engine

import "os"
import "io"
import "log"
import "fmt"
import "time"
import "sync"
import "path/filepath"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/mutations"
import "github.com/dadleyy/charlestown/engine/constants"
import "github.com/dadleyy/charlestown/engine/resources"

// Run creates and loops the game engine, mostly used to prepare the logger.
func Run(config Configuration) error {
	wd, e := os.Getwd()

	if e != nil {
		return e
	}

	if len(wd) == 1 && wd[0] == os.PathSeparator {
		wd = filepath.Join(os.Getenv("HOME"), ".charles-town")
	}

	logdir := filepath.Dir(config.Logging.Filename)

	if !filepath.IsAbs(logdir) {
		logdir = filepath.Join(wd, logdir)
	}

	if e := os.MkdirAll(logdir, os.ModePerm); e != nil {
		return fmt.Errorf("mkdirall fail on '%s' (working dir %s): %s", logdir, wd, e)
	}

	flags := os.O_RDWR | os.O_CREATE

	if config.Logging.Truncate {
		flags = flags | os.O_TRUNC
	} else {
		flags = flags | os.O_APPEND
	}

	logfile := filepath.Join(logdir, filepath.Base(config.Logging.Filename))
	writer, e := os.OpenFile(logfile, flags, os.ModePerm)

	if e != nil {
		return e
	}

	defer writer.Close()

	log.SetOutput(io.MultiWriter(writer))
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime | log.LUTC)

	world := objects.Dimensions{120, 40}

	cursor := objects.Cursor{
		Location: objects.Point{1, 1},
	}

	state := objects.Game{
		World:  world,
		Cursor: cursor,
		Funds:  5 * constants.EconomyMultiplier,
		Turn: objects.Turn{
			Actions: objects.TurnActions{0, 5 + constants.BaseTurnCardinality},
			Start:   time.Now(),
		},
	}

	log.Printf("[init] creating mutations channel")
	// Create a channel with our stream of mutations.
	updates := make(chan mutations.Mutation, 10)

	// Create a stop channel for our timing dispatcher.
	stop := make(chan struct{})
	wg := &sync.WaitGroup{}

	log.Printf("[init] creating timing dispatcher")
	// Create the timing dispatcher that will send updates to our engine.
	dispatch := &timingDispatcher{
		updates: updates,
		mutators: []mutator{
			&economyMutator{},
			&flashMutator{},
			&turnMutator{},
		},
	}

	log.Printf("[init] creating resource loader")
	loader, e := resources.NewFilesystemLoader(config.AssetPath)

	if e != nil {
		return e
	}

	log.Printf("[init] creating engine")
	instance := newOpenGLEngine(config, loader)

	log.Printf("[init] starting timing dispatcher")
	go dispatch.start(stop, wg)

	shutdown := instance.run(state, updates)

	log.Printf("[shutdown] sending stop signal to timers")
	stop <- struct{}{}

	log.Printf("[shutdown] closing update channel")
	close(updates)

	log.Printf("[shutdown] stop signal sent, waiting for wait group")
	wg.Wait()

	log.Printf("[shutdown] complete")
	return shutdown
}
