package engine

import "fmt"
import "log"
import "sync"
import "strings"
import "runtime"
import "github.com/go-gl/gl/v4.1-core/gl"
import "github.com/go-gl/glfw/v3.2/glfw"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"
import "github.com/dadleyy/charlestown/engine/mutations"
import "github.com/dadleyy/charlestown/engine/resources"

type openGLEngine struct {
	config Configuration
	loader resources.Loader
}

func (engine *openGLEngine) draw(window *glfw.Window, program uint32, game objects.Game) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	defer window.SwapBuffers()

	cx, cy := game.Cursor.Location.Values()
	wx, wy := game.World.Values()
	dx, dy := game.Dimensions.Values()

	if dx == 0 || dy == 0 || wx == 0 || wy == 0 {
		return nil
	}

	hry := float32(0.0)
	hrx := float32(0.0)

	// the ratio of our world to the viewport. If the world is wider, this will be > 1.0; smaller < 1.0.
	ax := float32(wx) / float32(dx)
	// map the ratio to the [-1, 1] space
	px := (ax * 2.0)

	// the ratio of our world to the viewport. If the world is taller, this will be > 1.0; smaller < 1.0.
	ay := float32(wy) / float32(dy)
	// map the ratio to [-1, 1] space
	py := (ay * 2.0)

	// origin x
	ox := -1.0 + float32(1.0/float32(dx))
	oy := 1.0 - float32(1.0/float32(dy))

	left := ox + hrx
	right := ox + px + hrx
	top := oy - hry
	bottom := oy - py - hry

	log.Printf("world (%d, %d) | viewport (%d, %d) | cursor (%d, %d) | bottom %.2f | right %.2f", wx, wy, dx, dy, cx, cy, bottom, right)

	borders := []float32{
		left, top, 0.0,
		right, top, 0.0,

		left, bottom, 0.0,
		right, bottom, 0.0,

		left, top, 0.0,
		left, bottom, 0.0,

		right, top, 0.0,
		right, bottom, 0.0,
	}

	gl.BindVertexArray(engine.makeVao(borders))
	gl.DrawArrays(gl.LINES, 0, int32(len(borders)/3))
	/*
			mx := (1.0 / (float32(dx) / 600))
			my := (1.0 / (float32(dy) / 400))
			px := float32(x) / float32(width)
			py := float32(y) / float32(height)

			// Draw the wall(s)
			borders := []float32{
				// top
				-px, py, 0.0,
				1.0 - px, py, 0.0,

				// bottom
				-px, py - 1.0, 0.0,
				1.0 - px, py - 1.0, 0.0,

				// left
				-px, py, 0.0,
				-px, py - 1.0, 0.0,

				// right
				1.0 - px, py, 0.0,
				1.0 - px, py - 1.0, 0.0,
			}
			gl.BindVertexArray(engine.makeVao(borders))
			gl.DrawArrays(gl.LINES, 0, int32(len(borders)/3))

			cursorWidth := float32((0.05 * mx) * 0.5)
			cursorHeight := float32((0.05 * my) * 0.5)

			gl.BindVertexArray(engine.makeVao([]float32{
				-cursorWidth, cursorHeight, 0.0,
				-cursorWidth, -cursorHeight, 0.0,
				cursorWidth, -cursorHeight, 0.0,

				-cursorWidth, cursorHeight, 0.0,
				cursorWidth, cursorHeight, 0.0,
				cursorWidth, -cursorHeight, 0.0,

					0.0, 0.0, 0.0,
					0.0, 0.0, 0.0,
					0.0, 0.0, 0.0,
			}))
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
	*/

	/*
		triangle := []float32{
			-0.5 + px, 0.5 - py, 0,
			-0.5 + px, -0.5 - py, 0,
			0.5 + px, -0.5 - py, 0,

			-0.5 + px, 0.5 - py, 0,
			0.5 + px, 0.5 - py, 0,
			0.5 + px, -0.5 - py, 0,
		}

		vao := engine.makeVao(triangle)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))
	*/

	return nil
}

func (engine *openGLEngine) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (engine *openGLEngine) makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (engine *openGLEngine) initWindow() (*glfw.Window, error) {
	if e := glfw.Init(); e != nil {
		log.Printf("[error] unable to initialize glfw: %s", e)
		return nil, e
	}

	monitors := glfw.GetMonitors()
	primary := glfw.GetPrimaryMonitor()

	for _, m := range monitors {
		x, y := m.GetPos()

		if primary == nil {
			primary = m
		}

		if px, _ := primary.GetPos(); x < px {
			primary = m
		}

		log.Printf("monitor (%d, %d): %s", x, y, m.GetName())
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	width, height := constants.DefaultWindowWidth, constants.DefaultWindowHeight
	window, e := glfw.CreateWindow(width, height, constants.WindowName, nil, nil)

	if primary != nil {
		x, y := primary.GetPos()
		window.SetPos(x, y)
	}

	if e != nil {
		return nil, e
	}

	window.MakeContextCurrent()

	return window, nil
}

func (engine *openGLEngine) initOpenGL() (uint32, error) {
	if e := gl.Init(); e != nil {
		return 0, e
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, e := engine.compileShader(constants.VertexShaderSource, gl.VERTEX_SHADER)
	if e != nil {
		return 0, e
	}

	fragmentShader, e := engine.compileShader(constants.FragmentShaderSource, gl.FRAGMENT_SHADER)
	if e != nil {
		return 0, e
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog, nil
}

// mutationForKey returns a mutation to be applied to the game based on a keyboard interaction.
func (engine *openGLEngine) mutationForKey(key glfw.Key, action glfw.Action) mutations.Mutation {
	mut := mutations.Move(0, 0)
	switch key {
	case glfw.KeyW:
		mut = mutations.Move(0, -10)
	case glfw.KeyA:
		mut = mutations.Move(-10, 0)
	case glfw.KeyS:
		mut = mutations.Move(0, 10)
	case glfw.KeyD:
		mut = mutations.Move(10, 0)
	case glfw.KeyTab:
		if action == glfw.Press {
			mut = mutations.Mode()
		}
	case glfw.KeyEnter:
		if action == glfw.Press {
			mut = mutations.Interact()
		}
	default:
		log.Printf("received unknown key movement %v", key)
	}
	return mut
}

// run initializes openl, glfw and registers the appropriate input callbacks + mutation loop.
func (engine *openGLEngine) run(game objects.Game, updates <-chan mutations.Mutation) error {
	runtime.LockOSThread()
	log.Printf("[init] starting glfw")
	window, e := engine.initWindow()

	if e != nil {
		log.Printf("[error] unable to initialize window: %s", e)
		return e
	}

	defer glfw.Terminate()

	log.Printf("[init] starting opengl")
	prog, e := engine.initOpenGL()

	if e != nil {
		log.Printf("[error] unable to initialize opengl: %s", e)
		return e
	}

	window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		modifier := (1.0 / (float32(width) / 600))
		log.Printf("received resize (%d x %d) [%.2f]", width, height, modifier)
		game.Dimensions = objects.Dimensions{width, height}
		engine.draw(window, prog, game)
	})

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		mut := engine.mutationForKey(key, action)
		game = mut.Apply(game)
		engine.draw(window, prog, game)
	})

	width, height := window.GetSize()
	log.Printf("[init] setting initial size (%d, %d)", width, height)
	game.Dimensions = objects.Dimensions{width, height}

	wg := &sync.WaitGroup{}
	quit := make(chan struct{})

	log.Printf("[init] starting update consumer")
	go func() {
		wg.Add(1)
		defer wg.Done()
		exit := false

		for exit == false {
			select {
			case update := <-updates:
				log.Printf("received update")
				game = update.Apply(game)
				glfw.PostEmptyEvent()
			case <-quit:
				exit = true
			}
		}

		log.Printf("[shutdown] finished update loop")
	}()

	log.Printf("[init] starting runloop")
	for !window.ShouldClose() {
		engine.draw(window, prog, game)
		glfw.WaitEvents()
	}

	quit <- struct{}{}
	log.Printf("[shutdown] runloop terminated")
	wg.Wait()
	return nil
}

func newOpenGLEngine(config Configuration, loader resources.Loader) engine {
	return &openGLEngine{loader: loader}
}
