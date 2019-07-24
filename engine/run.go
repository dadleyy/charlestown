package engine

import "github.com/gdamore/tcell"
import "github.com/gdamore/tcell/encoding"

// Run creates and loops the game engine.
func Run(config Configuration) error {
	logger, e := initializeLogger(config)

	if e != nil {
		return e
	}

	defer logger.Close()

	logger.Printf("[init] starting tcell")

	encoding.Register()

	screen, e := tcell.NewScreen()

	if e != nil {
		return e
	}

	logger.Printf("[init] tcell started, entering game loop")

	instance := engine{logger, screen, config}
	return instance.run()
}
