package main

//go:generate go mod init github.com/donomii/splash-screen
//go:generate go mod tidy

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
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
	VertAttrib     int32
	ColourAttrib   int32
	Angle          float64
	PreviousTime   float64
	ModelUniform   int32
	TexCoordAttrib int32
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
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	win, err := glfw.CreateWindow(winWidth, winHeight, "Lsystems", nil, nil)
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

	win.SetMouseButtonCallback(handleMouseButton)

	win.SetCursorPosCallback(handleMouseMove)

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

	state.Vao, state.Vbo, state.VertAttrib = make_array_buffer("vert", 3, state.Program, gl.FLOAT)
	state.Cao, state.Cbo, state.ColourAttrib = make_array_buffer("s_col", 4, state.Program, gl.FLOAT)
	gl.BindFragDataLocation(state.Program, 0, gl.Str("outputColor\x00"))
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

	now := glfw.GetTime()
	elapsed := now - state.PreviousTime

	// Configure global settings

	gl.UseProgram(state.Program)
	gl.ClearColor(0.5, 0.5, 0.5, 1.0)

	gl.Disable(gl.BLEND)
	gl.Enable(gl.DEPTH_TEST)

	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	if elapsed > 0.050 && 1 != win.GetAttrib(glfw.Iconified) {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		angle := state.Angle
		angle += elapsed
		state.Angle = angle

		vertices, colours := calcLsys()

		//fmt.Printf("array len: %v\n", len(colours))

		gl.BindVertexArray(state.Vao)
		checkGlError()
		gl.BindBuffer(gl.ARRAY_BUFFER, state.Vbo)
		checkGlError()
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
		checkGlError()
		gl.VertexAttribPointer(uint32(state.VertAttrib), 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
		checkGlError()
		gl.EnableVertexAttribArray(uint32(state.VertAttrib))
		checkGlError()

		gl.BindBuffer(gl.ARRAY_BUFFER, state.Cbo)
		checkGlError()
		gl.BufferData(gl.ARRAY_BUFFER, len(colours)*4, gl.Ptr(colours), gl.STATIC_DRAW)
		checkGlError()
		gl.VertexAttribPointer(uint32(state.ColourAttrib), 4, gl.FLOAT, false, 0, gl.PtrOffset(0))
		checkGlError()
		gl.EnableVertexAttribArray(uint32(state.ColourAttrib))
		checkGlError()

		gl.BindVertexArray(state.Vao)
		checkGlError()
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/3))
		checkGlError()

		win.SwapBuffers()

	}
	time.Sleep(10 * time.Millisecond)
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

type immob struct {
	Location []float32
	Lsys     string
}

var sceneList []immob
var scale float32 = 1

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

	scale = 10.0 //scale + 1.0
	//fmt.Printf("scale: %v\n", scale)

	/*
		objs := []string{
			" HR s s s s s s s s leaf ",
			" HR s s s s s Arrow ",
			" HR s s s s s s s s Flower ",
			" HR s s s s s s s s Circle ",
			" HR s s s s s s s s Icosahedron ",
		}

			sceneList = []immob{}
			for i := 0; i < 5; i++ {
				imm := immob{[]float32{float32(i) / scale, float32(i) / scale, float32(i) / scale}, objs[i]}
				sceneList = append(sceneList, imm)

			}
	*/

	for _, imm := range sceneList {
		movMatrix := mgl32.Ident4()
		movMatrix = Move(movMatrix, imm.Location[0], imm.Location[1], imm.Location[2])
		//fmt.Printf("moatrix: %v\n", movMatrix)
		a, b := lsystem.Draw(CurrentScene, scene_camera.ViewMatrix(), lsystem.S(imm.Lsys), movMatrix, true)
		verticesNative = append(verticesNative, a...)
		colorsNative = append(colorsNative, b...)
	}

	movMatrix = mgl32.Ident4()
	movMatrix = Move(movMatrix, x, y, 0)
	//fmt.Printf("moatrix: %v\n", movMatrix)
	a, b := lsystem.Draw(CurrentScene, scene_camera.ViewMatrix(), lsystem.S(" HR s s s s s Arrow "), movMatrix, true)
	verticesNative = append(verticesNative, a...)
	colorsNative = append(colorsNative, b...)

	return verticesNative, colorsNative
}

var x, y float32

func handleMouseMove(win *glfw.Window, xpos float64, ypos float64) {
	w, h := win.GetSize()
	x = float32((xpos-float64(w)/2)/float64(w)) * 2
	y = -float32((ypos-float64(h)/2)/float64(h)) * 2

}

func handleMouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Got mouse button %v,%v,%v", button, mod, action)
	fmt.Println("Click at", x, y)
	imm := immob{[]float32{x, y, rand.Float32()}, " HR s s s s s Arrow "}
	sceneList = append(sceneList, imm)
	//handleKey(w, key, scancode, action, mods)
}
