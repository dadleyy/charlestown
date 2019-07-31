package engine

import "os"
import "io"
import "log"
import "path/filepath"
import "github.com/dadleyy/charlestown/engine/objects"

// Run creates and loops the game engine, mostly used to prepare the logger.
func Run(config Configuration) error {
	wd, e := os.Getwd()

	if e != nil {
		return e
	}

	logdir := filepath.Dir(config.Logging.Filename)

	if !filepath.IsAbs(logdir) {
		logdir = filepath.Join(wd, logdir)
	}

	if e := os.MkdirAll(logdir, os.ModePerm); e != nil {
		return e
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

	logger := log.New(io.MultiWriter(writer), "[ch] ", log.Ldate|log.Lshortfile|log.Ltime|log.LUTC)

	world := objects.Dimensions{120, 40}

	cursor := objects.Cursor{
		Location: objects.Point{world.Width / 2, world.Height / 4},
	}
	state := objects.Game{
		World:  world,
		Cursor: cursor,
	}

	instance := engine{
		Logger: logger,
		config: config,
		renderers: []renderer{
			&boundaryRenderer{logger},
			&buildingRenderer{logger},
			&cursorRenderer{logger},
			&uiRenderer{logger, true},
		},
	}

	return instance.run(state)
}
