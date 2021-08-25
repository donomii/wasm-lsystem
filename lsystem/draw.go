package lsystem

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"
)

var green float32

func reverse3(t []float32) {
	s := len(t) - 6
	a := mgl32.Vec3{t[s], t[s+1], t[s+2]}
	b := mgl32.Vec3{t[s+3], t[s+4], t[s+5]}

	t[s], t[s+1], t[s+2] = b[0], b[1], b[2]
	t[s+3], t[s+4], t[s+5] = a[0], a[1], a[2]
}

func compose(a, b mgl32.Mat4) mgl32.Mat4 {
	return a.Mul4(b)
}
func pushState(aStack []mgl32.Mat4, aVal mgl32.Mat4) []mgl32.Mat4 {
	return append(aStack, aVal)
}

func popState(aStack []mgl32.Mat4, aVal mgl32.Mat4) ([]mgl32.Mat4, mgl32.Mat4) {
	//fmt.Printf("Stack length: %f\n", len(aStack))
	//fmt.Printf("Stack : %v\n", aStack)
	if len(aStack) < 1 {
		fmt.Println("Pop called on empty stack!")
		//return []mgl32.Mat4{}, aVal
		os.Exit(0)
	}
	if len(aStack) == 1 {
		return []mgl32.Mat4{}, aStack[0]
	}
	if len(aStack) == 2 {
		return []mgl32.Mat4{aStack[0]}, aStack[1]
	}
	if len(aStack) == 3 {
		return []mgl32.Mat4{aStack[0], aStack[1]}, aStack[2]
	}
	//if (len(aStack) == 3 ) {
	//	return aStack[0:1], aStack[2]
	//}
	return aStack[:len(aStack)-1], aStack[len(aStack)-1]
}

func pushAttribs(aStack []attribs, aVal attribs) []attribs {
	return append(aStack, aVal)
}

func popAttribs(aStack []attribs, aVal attribs) ([]attribs, attribs) {
	//fmt.Printf("Stack length: %f\n", len(aStack))
	//fmt.Printf("Stack : %v\n", aStack)
	if len(aStack) < 1 {
		//fmt.Println("Pop called on empty stack!")
		return []attribs{}, aVal
		os.Exit(0)
	}
	if len(aStack) == 1 {
		return []attribs{}, aStack[0]
	}
	if len(aStack) == 2 {
		return []attribs{aStack[0]}, aStack[1]
	}
	if len(aStack) == 3 {
		return []attribs{aStack[0], aStack[1]}, aStack[2]
	}
	//if (len(aStack) == 3 ) {
	//	return aStack[0:1], aStack[2]
	//}
	return aStack[:len(aStack)-1], aStack[len(aStack)-1]
}

func quad() []float32 {
	var triangleData = []float32{
		-1.0, 1.0, 0.0, // top left
		-1.0, -1.0, 0.0, // bottom left
		1.0, 1.0, 0.0, // bottom right
		-1.0, -1.0, 0.0, // bottom right
		1.0, -1.0, 0.0, // top left
		1.0, 1.0, 0.0, // bottom right
	}
	return triangleData
}
func checkGlErr() {}

var setAngleRegex = regexp.MustCompile(`A([0-9.]+)`)
var setDegreeRegex = regexp.MustCompile(`deg([0-9.]+)`)
var setHingeRegex = regexp.MustCompile(`Hinge\(([0-9]+)\)`)
var setColourRegex = regexp.MustCompile(`Colour(\d\d?\d?),(\d\d?\d?),(\d\d?\d?)`)
var setScaleRegex = regexp.MustCompile(`Scale\((-?[0-9.]+),(-?[0-9.]+),(-?[0-9.]+)\)`)

func Draw(CurrentScene *Scene, camera mgl32.Mat4, start []string, trans mgl32.Mat4, buildMode bool) ([]float32, []float32) {

	//fmt.Printf("Start: %v\n", start)
	triBuf := []float32{}
	colBuf := []float32{}
	commands := start
	if buildMode {
		//fmt.Println(ruleBook())
		commands = runRuleset(start, ruleBook())
		//trans = mgl32.Ident4()
	}
	//commands = runRules(commands, rules(), 2)

	forward := []float32{0, 1.0, 0, 0}
	//PI := float32(3.1415927)
	a := attribs{angle: float32(0.2), red: 1.0, blue: 1.0, green: 1.0, alpha: 1.0, mirror: false}
	stateStack := []mgl32.Mat4{}
	//fmt.Printf("Commands: %v\n", commands)

	for _, c := range commands {
		//fmt.Printf("Processing command: '%v'\n", c)
		//fmt.Println("Angle: ",angle)

		//fmt.Println(trans)
		switch {
		case c == "":
		case c == "stem":
			trans = compose(trans, mgl32.Scale3D(1.0, 1.0, 1.0))
		case c == "F":
			trans = compose(trans, mgl32.Translate3D(forward[0], forward[1], forward[2]))
		case c == "f":
			trans = compose(trans, mgl32.Translate3D(-forward[0], -forward[1], -forward[2]))
		case c == "Y":
			trans = compose(trans, mgl32.HomogRotate3DZ(a.angle))
		case c == "R":
			trans = compose(trans, mgl32.HomogRotate3DY(a.angle))
		case c == "P":
			trans = compose(trans, mgl32.HomogRotate3DX(a.angle))
		case c == "y":
			trans = compose(trans, mgl32.HomogRotate3DZ(-a.angle))
		case c == "r":
			trans = compose(trans, mgl32.HomogRotate3DY(-a.angle))
		case c == "p":
			trans = compose(trans, mgl32.HomogRotate3DX(-a.angle))
		case c == "hs":
			trans = compose(trans, mgl32.Scale3D(0.5, 0.5, 0.5))
		case c == "hS":
			trans = compose(trans, mgl32.Scale3D(2.0, 2.0, 2.0))
		case c == "s":
			trans = compose(trans, mgl32.Scale3D(0.666, 0.666, 0.666))
		case c == "S":
			trans = compose(trans, mgl32.Scale3D(1.5, 1.5, 1.5))
		case c == "SY":
			trans = compose(trans, mgl32.Scale3D(1.0, 2.0, 1.0))
		case c == "[":
			stateStack = pushState(stateStack, trans)
			attribStack = pushAttribs(attribStack, a)
			//log.Printf("Push!\n")
		case c == "]":
			stateStack, trans = popState(stateStack, trans)
			attribStack, a = popAttribs(attribStack, a)
			//log.Printf("Pop!\n")
		case setHingeRegex.FindString(c) != "":
			var match = setHingeRegex.FindStringSubmatch(c)
			parsedNum, _ := strconv.ParseInt(match[1], 10, 32)
			num := int(parsedNum)
			trans = compose(trans, mgl32.HomogRotate3DY(hinges[num]))
		case c == "HR":
			trans = compose(trans, mgl32.HomogRotate3DY(CurrentScene.Clock*3.14159*2.0))
		case c == "rotateGreen":
			green += 0.05
			if green > 1 {
				green = 0.8
			}
		case setColourRegex.FindString(c) != "":
			////fmt.Println(c)
			var match = setColourRegex.FindStringSubmatch(c)
			parsedNum, _ := strconv.ParseFloat(match[1], 32)
			a.red = float32(parsedNum) / 255
			parsedNum, _ = strconv.ParseFloat(match[2], 32)
			a.green = float32(parsedNum / 255)
			parsedNum, _ = strconv.ParseFloat(match[3], 32)
			a.blue = float32(parsedNum / 255)
			a.alpha = 1.0
		case setAngleRegex.FindString(c) != "":
			var match = setAngleRegex.FindStringSubmatch(c)
			parsedNum, _ := strconv.ParseFloat(match[1], 32)
			a.angle = float32(parsedNum)
		case setDegreeRegex.FindString(c) != "":
			var match = setDegreeRegex.FindStringSubmatch(c)
			parsedNum, _ := strconv.ParseFloat(match[1], 32)
			a.angle = float32(parsedNum / 180.0 * math.Pi)
		case setScaleRegex.FindString(c) != "":
			var match = setScaleRegex.FindStringSubmatch(c)
			parsedNum, _ := strconv.ParseFloat(match[1], 32)
			x := float32(parsedNum)
			parsedNum, _ = strconv.ParseFloat(match[2], 32)
			y := float32(parsedNum)
			parsedNum, _ = strconv.ParseFloat(match[3], 32)
			z := float32(parsedNum)
			//fmt.Println("Scaling by ", x, y, z)
			trans = compose(trans, mgl32.Scale3D(x, y, z))
		case c == "reverseTriangle":
			reverse3(triBuf)
		case c == ".":
			//idVec := mgl32.Vec4{1,0,0,0}
			//newPoint := trans.Mul4x1(idVec)
			//triBuf = append(triBuf, newPoint[0], newPoint[1], newPoint[2])
			if buildMode {
				triBuf = append(triBuf, point(trans)...)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}
		case c == "op":
			if buildMode {
				triBuf = append(triBuf, point(trans)...)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}

		case c == "origin":
			if buildMode {
				idVec := mgl32.Vec4{0, 0, 0, 0}
				triBuf = append(triBuf, idVec[0], idVec[1], idVec[2])
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}
		case c == "T":
			if buildMode {
				//glctx.Uniform4f(color, a.red, a.green, a.blue, 1)
				triBuf = append(triBuf, tr(trans)...)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}
		case c == "Q":
			if buildMode {
				//glctx.Uniform4f(color, a.red, a.green, a.blue, 1)
				log.Printf("Pushing quad vertices\n")
				triBuf = append(triBuf, quad()...)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}
		case c == "TF":
			if buildMode {
				triBuf = append(triBuf, tr(trans)...)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
				colBuf = append(colBuf, a.red, a.green, a.blue, a.alpha)
			}
			trans = compose(trans, mgl32.Translate3D(forward[0], forward[1], forward[2]))
		case c == "LightsOn":
			a.useLighting = true
		case c == "LightsOff":
			a.useLighting = false
		default:
			//log.Println("'", c, "'")

			//log.Println("Paint: ", c)
			//paintCube(camera, trans, c, a, ModelMatrix, gl, indicesNative)

		}
	}
	return triBuf, colBuf
}

/*
func paintCube(camera mgl32.Mat4, trans mgl32.Mat4, name string, a attribs, ModelMatrix js.Value, gl js.Value, indicesNative []uint16) {
	var glTypes gltypes.GLTypes
	glTypes.New(gl)
	if len(name) == 0 {
		//log.Printf("Empty name, refusing to draw")
		return
	}
	log.Printf("Starting %v", name)

	checkGlErr()

	log.Println("Drawing object:", name)
	final := trans.Mul4(mgl32.Scale3D(0.5, 0.5, 0.5))
	//fmt.Println(final)
	// Convert model matrix to a JS TypedArray
	var modelMatrixBuffer *[16]float32
	modelMatrixBuffer = (*[16]float32)(unsafe.Pointer(&final))
	typedModelMatrixBuffer := gltypes.SliceToTypedArray([]float32((*modelMatrixBuffer)[:]))
	// Apply the model matrix
	gl.Call("uniformMatrix4fv", ModelMatrix, false, typedModelMatrixBuffer)
	// Draw the cube
	//gl.Call("drawElements", glTypes.Triangles, len(indicesNative), glTypes.UnsignedShort, 0)

	checkGlErr()

}
*/
/*
func paintVertices(camera mgl32.Mat4, trans mgl32.Mat4, name string, a attribs, ModelMatrix js.Value, gl js.Value) {
	var glTypes gltypes.GLTypes
	glTypes.New(gl)
	if len(name) == 0 {
		//log.Printf("Empty name, refusing to draw")
		return
	}
	log.Printf("Starting %v", name)

	checkGlErr()

	log.Println("Drawing object:", name)
	final := trans.Mul4(mgl32.Scale3D(0.5, 0.5, 0.5))
	//fmt.Println(final)
	// Convert model matrix to a JS TypedArray
	var modelMatrixBuffer *[16]float32
	modelMatrixBuffer = (*[16]float32)(unsafe.Pointer(&final))
	typedModelMatrixBuffer := gltypes.SliceToTypedArray([]float32((*modelMatrixBuffer)[:]))
	// Apply the model matrix
	gl.Call("uniformMatrix4fv", ModelMatrix, false, typedModelMatrixBuffer)
	// Draw the cube
	//gl.Call("drawElements", glTypes.Triangles, len(indicesNative), glTypes.UnsignedShort, 0)

	checkGlErr()
}
*/
