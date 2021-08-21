package main

//go:generate go mod init github.com/donomii/splash-screen
//go:generate go mod tidy

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	//"time"

	"../lsystem"
	"github.com/donomii/sceneCamera"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"

	_ "embed"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
	debug.SetGCPercent(-1)
}

type State struct {
	prop           int32
	Program        uint32
	Vao            uint32
	Vbo            uint32
	Cao            uint32
	Cbo            uint32
	Texture        uint32
	TextureUniform int32
	VertAttrib     uint32
	ColourAttrib   uint32
	Angle          float64
	PreviousTime   float64
	ModelUniform   int32
	TexCoordAttrib uint32
}

var winWidth = 180
var winHeight = 180
var lasttime float64

var rotX, roty float64
var CurrentScene *lsystem.Scene
var scene_camera *sceneCamera.SceneCamera = sceneCamera.New()

func drainChannel(ch chan []byte) {
	for {
		<-ch
	}
}
func main() {
	flag.Parse()

	currentDir, _ := os.Getwd()

	os.Chdir(currentDir)
	log.Println("Starting windowing system")
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Grafana", nil, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			CurrentScene.Clock += 0.01
		}
	}()

	win.MakeContextCurrent()

	win.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		log.Printf("Got key %c,%v,%v,%v", key, key, mods, action)
		//handleKey(w, key, scancode, action, mods)
	})

	//win.SetMouseButtonCallback(handleMouseButton)

	//win.SetCursorPosCallback(handleMouseMove)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	state := &State{
		prop: 1,
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	state.Program, err = newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	//Activate the program we just created.  This means we will use the render and fragment shaders we compiled above
	gl.UseProgram(state.Program)

	//Set a default projection matrix
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(winWidth)/float32(winHeight), 0.1, 10.0)
	projectionUniform := gl.GetUniformLocation(state.Program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	//Setup the camera
	camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraUniform := gl.GetUniformLocation(state.Program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	//Setup the cube
	model := mgl32.Ident4()
	state.ModelUniform = gl.GetUniformLocation(state.Program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(state.ModelUniform, 1, false, &model[0])

	//Find the location of the texture, so we can upload a picture to it
	state.TextureUniform = gl.GetUniformLocation(state.Program, gl.Str("tex\x00"))
	gl.Uniform1i(state.TextureUniform, 0)

	//This is the variable in the fragment shader that will hold the colour for each pixel
	gl.BindFragDataLocation(state.Program, 0, gl.Str("outputColor\x00"))

	// Configure the vertex data
	gl.GenVertexArrays(1, &state.Vao)
	gl.BindVertexArray(state.Vao)

	gl.GenBuffers(1, &state.Vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, state.Vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
	checkGlError()
	state.VertAttrib = uint32(gl.GetAttribLocation(state.Program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(state.VertAttrib)
	gl.VertexAttribPointer(state.VertAttrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	checkGlError()

	gl.GenVertexArrays(1, &state.Cao)
	gl.BindVertexArray(state.Cao)

	gl.GenBuffers(1, &state.Cbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, state.Cbo)
	state.ColourAttrib = uint32(gl.GetAttribLocation(state.Program, gl.Str("s_col\x00")))
	gl.EnableVertexAttribArray(state.VertAttrib)
	gl.VertexAttribPointer(state.VertAttrib, 4, gl.FLOAT, false, 0, gl.PtrOffset(0))
	checkGlError()

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.UseProgram(state.Program)
	gl.ClearColor(1.0, 1.0, 1.0, 0.0)

	//Activate the cube data, which will be drawn
	gl.BindVertexArray(state.Vao)

	//Choose the texture we just created and uploaded
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, state.Texture)

	resetCam(scene_camera)
	sceneList := lsystem.InitScenes(scene_camera)
	CurrentScene = sceneList[0]
	CurrentScene.Init(CurrentScene)

	for !win.ShouldClose() {

		gfxMain(win, state)
		glfw.PollEvents()
	}
	shutdown()

}

func shutdown() {

}

func gfxMain(win *glfw.Window, state *State) {
	//fmt.Println("Draw")
	//width, height := win.GetSize()
	//gl.Viewport(0, 0, int32(width-1), int32(height-1))

	// Render

	// Update

	now := glfw.GetTime()
	elapsed := now - state.PreviousTime

	if elapsed > 0.050 && 1 != win.GetAttrib(glfw.Iconified) {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		//fmt.Printf("elapsed: %v\n", elapsed)
		state.PreviousTime = now
		angle := state.Angle
		angle += elapsed
		state.Angle = angle

		//model := mgl32.HomogRotate3D(float32(angle+rotX), mgl32.Vec3{0, 1, 0})

		// Render

		vertices, colours := calcLsys()
		fmt.Printf("array: %v\n", colours)

		gl.BindBuffer(gl.ARRAY_BUFFER, state.ColourAttrib)
		gl.BufferData(gl.ARRAY_BUFFER, len(colours)*4, gl.Ptr(colours), gl.STATIC_DRAW)
		//gl.UniformMatrix4fv(state.ModelUniform, 1, false, &vertices[0])
		gl.BindBuffer(gl.ARRAY_BUFFER, state.VertAttrib)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)))

		win.SwapBuffers()

	}
	time.Sleep(10 * time.Millisecond)
}

func checkGlError() {

	err := gl.GetError()
	if err > 0 {
		errStr := fmt.Sprintf("GLerror: %v\n", err)
		fmt.Printf(errStr)
		panic(errStr)
	}

}

func Move(movMatrix mgl32.Mat4, x, y, z float32) mgl32.Mat4 {
	movMatrix = movMatrix.Mul4(mgl32.Translate3D(x, y, z))
	return movMatrix
}
func resetCam(camera *sceneCamera.SceneCamera) {
	camera.Reset()
	camera.SetPosition(0, 0, 1)
	camera.Translate(0.0, 0.0, 0.0)
	camera.LookAt(0.0, 0.0, 0.0)
}
func calcLsys() ([]float32, []float32) {

	movMatrix := mgl32.Ident4()
	movMatrix = Move(movMatrix, 1.0, 0.0, 0.0)

	verticesNative, colorsNative := lsystem.Draw(CurrentScene, scene_camera.ViewMatrix(), lsystem.S(`
			Colour254,254,254 deg30 s s s s  r r F p p p f f f f [ HR
				s s s s
				[ s s HR Icosahedron ] TF TF TF TF 
				[ HR Tetrahedron ] Arrow  F  Arrow  F  Arrow  F  
				[ p p p s s s HR starburst ] Arrow  F  Arrow  F  Arrow  F 
				[ p p p s s HR leaf ] Arrow  F  Arrow  F  Arrow  F 	
				
				[ p p p s s s HR lineStar ] TF TF TF
				[ p p p s s HR Flower ] TF TF TF
				[ p p p s s HR Flower12 ] TF TF TF
				[ p p p s s HR Flower11 ] TF TF TF
				[ p p p s s HR Flower10 ] TF TF TF
				
				
			]
			
			p p p F P P P
			[ s s s s
			
				
				[ p p p S S S HR Square1 ] TF TF TF
				[ p p p S S S S S S HR Face ] TF TF TF
				[ p p p S S S HR Arrow ] TF TF TF
				[ p p p S HR Prism ] TF TF TF
				[ p p p S HR Prism1 ] TF TF TF
				[   s s HR p p p Circle ] TF TF TF
				
				
					
				
				
			
				
			]
			
			`), movMatrix, true)

	return verticesNative, colorsNative
}
