package engine

import "log"
import "github.com/dadleyy/charlestown/engine/objects"

type openGLEngine struct {
}

func (engine *openGLEngine) run(state objects.Game) error {
	return nil
}

func newOpenGLEngine(config Configuration) engine {
	log.Printf("initializing opengl")
	return &openGLEngine{}
}
