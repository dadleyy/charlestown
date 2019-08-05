package engine

import "fmt"
import "log"
import "strings"
import "runtime"
import "github.com/go-gl/gl/v4.1-core/gl"
import "github.com/go-gl/glfw/v3.2/glfw"
import "github.com/dadleyy/charlestown/engine/objects"
import "github.com/dadleyy/charlestown/engine/constants"

type openGLEngine struct {
}

func (engine *openGLEngine) draw(window *glfw.Window, program uint32) error {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	triangle := []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}

	vao := engine.makeVao(triangle)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
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
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	width, height := constants.DefaultWindowWidth, constants.DefaultWindowHeight
	window, e := glfw.CreateWindow(width, height, constants.WindowName, nil, nil)

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

func (engine *openGLEngine) run(state objects.Game) error {
	runtime.LockOSThread()

	log.Printf("[init] starting glfw")
	window, e := engine.initWindow()

	if e != nil {
		return e
	}

	defer glfw.Terminate()

	prog, e := engine.initOpenGL()

	if e != nil {
		return e
	}

	log.Printf("[init] starting runloop")
	for !window.ShouldClose() {
		engine.draw(window, prog)
	}

	return nil
}

func newOpenGLEngine(config Configuration) engine {
	return &openGLEngine{}
}
