package lsystem

import (
	"regexp"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
)

var polyCount int           //Internal tracker to report the polygons used per scene
var compiled []string = nil //Cache completed lsystems
var attribStack []attribs   //The should be local, maybe

const ( // iota is reset to 0
	TRIANGLES    = iota //  0
	TRIANGLE_FAN = iota //  1
	LINES        = iota
	LINE_LOOP    = iota // 2
	LINE_STRIP   = iota
)

type ruleMeta struct {
	iterations   int
	x            float32
	y            float32
	style        int
	backfaceCull bool
	useLighting  bool
}

type attribs struct {
	angle       float32
	red         float32
	green       float32
	blue        float32
	alpha       float32
	useLighting bool
	mirror      bool
}

//Splitstring
func s(str string) []string {
	str = regexp.MustCompile("\n").ReplaceAllLiteralString(str, " ")
	str = regexp.MustCompile("\r").ReplaceAllLiteralString(str, " ")
	str = regexp.MustCompile("\t").ReplaceAllLiteralString(str, " ")
	str = regexp.MustCompile("  ").ReplaceAllLiteralString(str, " ")
	str = regexp.MustCompile("  ").ReplaceAllLiteralString(str, " ")
	return strings.Split(str, " ")
}

func rewrite(aString []string, rules map[string][]string) []string {

	//fmt.Println(aString)
	var outString []string
	for _, orig := range aString {
		new, ok := rules[orig]
		//fmt.Println(orig, new)
		if !ok {
			outString = append(outString, []string{orig}...)
		} else {
			outString = append(outString, new...)
		}
	}
	return outString
}

func ruleBook() map[string]map[string][]string {
	rb := map[string]map[string][]string{
		"ilake": map[string][]string{
			"ilake":      strings.Split("A1.57 ilakeiter Y ilakeiter Y ilakeiter Y ilakeiter", " "),
			"ilakeiter":  strings.Split("A1.57 ilakeiter Y ilakeiterf y ilakeiter ilakeiter Y ilakeiter Y ilakeiter ilakeiter Y ilakeiter ilakeiterf Y ilakeiter ilakeiter y ilakeiterf Y ilakeiter ilakeiter y ilakeiter y ilakeiter ilakeiterf y ilakeiter ilakeiter ilakeiter", " "),
			"ilakeiterf": strings.Split("TF ilakeiterf ilakeiterf ilakeiterf ilakeiterf ilakeiterf ilakeiterf", " "),
		},
		"Quad": map[string][]string{
			"Quad": strings.Split("Q", " "),
		},
		"Koch3": map[string][]string{
			"Koch3": strings.Split("R A1.57 TF y TF y TF y TF", " "),
			"TF":    strings.Split("TF TF y TF y TF y TF y TF y TF Y TF", " "),
		},
		"Tree3": map[string][]string{
			"Tree3": strings.Split("A0.3926991 TF", " "),
			"TF":    strings.Split("TF TF y [ y TF Y TF Y TF ] Y [ Y TF y TF y TF ]", " "),
		},
		"3DTree3": map[string][]string{
			"3DTree3": strings.Split("Colour255,233,155 A0.3926991 TF", " "),
			"TF":      strings.Split("TF TF y r [ y TF Y TF Y TF ] Y [ Y TF y TF y TF ]", " "),
		},
		"3DTreeLeafy": map[string][]string{
			"3DTreeLeafy": strings.Split("Colour0,255,255 A0.3926991 TF", " "),
			"TF":          strings.Split("TF TF y r le Colour0,255,255 [ y TF Y TF Y TF [ s s s s s s leaf ] Colour0,255,255  ] Y [ Y TF y TF y TF [ s s s s s s leaf ] Colour0,255,255 ]  Colour0,255,255", " "),
		},
		"starburst": map[string][]string{
			"starburst": strings.Split("A0.35 [ lA ] [ lB ]", " "),
			"lA":        strings.Split("[ y lA ] lC", " "),
			"lB":        strings.Split("[ Y lB ] lC", " "),
			"lC":        strings.Split("TF lC", " "),
		},
		"leaf": map[string][]string{
			"leaf": strings.Split("A0.35 Colour0,255,0 [ lA ] [ lB ]", " "),
			"lB":   strings.Split("[ Y lB { Colour0,155,0 . ] Colour0,255,0 . lC Colour0,155,0 . }", " "),
			"lA":   strings.Split("[ y lA { Colour0,155,0 . ] Colour0,155,0 . lC Colour0,255,0 . reverseTriangle }", " "),
			"lC":   strings.Split("F lC", " "),
		},
		"lineStar": map[string][]string{
			"lineStar": strings.Split("A0.35 Colour0,255,0 [ lA ] [ lB ]", " "),
			"lA":       strings.Split("[ y lA { . ] . lC . }", " "),
			"lB":       strings.Split("[ Y lB { . ] . lC . }", " "),
			"lC":       strings.Split("F lC", " "),
		},
		"orient": map[string][]string{
			"orient": strings.Split("TF A1.50 Colour255,0,0 starburst Y Y starburst P F F F Colour0,255,0 starburst Y Y starburst", " "),
		},
		"leaf2": map[string][]string{
			"leaf2": strings.Split("A0.35 Colour0,255,0 . [ lA ] [ lB ]", " "),
			"lA":    strings.Split("[ y lA ] lC .", " "),
			"lB":    strings.Split("[ Y lB ] lC .", " "),
			"lC":    strings.Split("F lC", " "),
		},
		"risingVine": map[string][]string{
			"risingVine": strings.Split("[ A0.35 AA ] A1.57 r r y y [ A0.35 AA ]", " "),
			"AA":         strings.Split("TF TF TF TF s AA", " "),
		},
		"KIomega": map[string][]string{
			"KIomega": strings.Split("A1.57 s s s s s TF y TF y TF y TF y", " "),
			"TF":      strings.Split("TF y TF Y TF Y TF TF y TF y TF Y TF", " "),
		},
		"Gosper": map[string][]string{
			"Gosper":  strings.Split("Colour255,255,255 A1.0472 Gosperl", " "),
			"Gosperl": strings.Split(". F Gosperl Y Gosperr Y Y Gosperr y Gosperl y y Gosperl Gosperl y Gosperr Y", " "),
			"Gosperr": strings.Split(". F y Gosperrlr Y Gosperr Gosperr Y Y Gosperr Y Gosperl y y Gosperl y Gosperr", " "),
		},
		"Sierpinksi": map[string][]string{
			"Sierpinksi": strings.Split("A1.0472 Sierpr", " "),
			"Sierpl":     strings.Split(". F Sierpr Y Sierpl Y Sierpr", " "),
			"Sierpr":     strings.Split(". F Sierpl y Sierpr y Sierpl", " "),
		},
		"Koch2": map[string][]string{
			"Koch2": strings.Split("A1.57 y TF", " "),
			"TF":    strings.Split("TF Y TF y TF y TF Y TF T F", " "),
		},
		"Square": map[string][]string{
			"Square": strings.Split("Colour0,0,255 A1.57 TF y y TF", " "),
		},
		"Plant": map[string][]string{
			"Plant":     strings.Split("A0.314159 Colour139,69,19 Plant1", " "),
			"Plant1":    strings.Split("A0.314159 Colour139,69,19 internode Y [ Plant1 + Flower ] y y R R [ y y s s s leaf ] internode [ Y Y s s s leaf ] y [ Plant1 Flower ] Y Y Plant1 Flower", " "),
			"internode": strings.Split("TF seg [ R R p p s s leaf ] [ R R P P s s leaf ] TF seg", " "),
			"seg":       strings.Split("seg TF seg", " "),
		},
		"Flower": map[string][]string{
			"Flower":  strings.Split("A0.314159 Colour139,69,19 [ Pedicel RotateColour Y Wedge Y Y Y Y Wedge Y Y Y Y Wedge Y Y Y Y Wedge Y Y Y Y Wedge ]", " "),
			"Pedicel": strings.Split("TF TF", " "),
			"Wedge":   strings.Split("[ Colour0,255,0 P TF ] [ WedgeLeaf ]", " "),
		},
		"Flower12": map[string][]string{
			"Flower12": strings.Split("A0.523598 Colour139,69,19 [ Pedicel RotateColour Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge  ]", " "),
			"Pedicel":  strings.Split("TF TF", " "),
			"Wedge":    strings.Split("[ Colour0,255,0 P TF ] [ WedgeLeaf ]", " "),
		},
		"Flower11": map[string][]string{
			"Flower11": strings.Split("A0.523598 Colour139,69,19 [ Pedicel RotateColour Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge ]", " "),
			"Pedicel":  strings.Split("TF TF", " "),
			"Wedge":    strings.Split("[ Colour0,255,0 P TF ] [ WedgeLeaf ]", " "),
		},
		"Flower10": map[string][]string{
			"Flower10": strings.Split("A0.523598 Colour139,69,19 [ Pedicel RotateColour Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge Y Wedge ]", " "),
			"Pedicel":  strings.Split("TF TF", " "),
			"Wedge":    strings.Split("[ Colour0,255,0 P TF ] [ WedgeLeaf ]", " "),
		},
		"FlowerField": map[string][]string{
			"FlowerField": strings.Split("Flower A1.57 s s P y [ spiral ]", " "),
			"column":      strings.Split("column [ A0.75 Y A1.5708 [ row A1.57 ] ] F F F", " "),
			"columnb":     strings.Split("columnb [ A0.75 Y A1.5708 [ row A1.57 ] ] f f f", " "),
			"row":         strings.Split("row A0.2 p F [ Flower ]", " "),
			"spiral":      strings.Split("TF [ Flower ] A0.2 Y Scale(0.9,0.9,0.9) spiral", " "),
		},
		"WedgeLeaf": map[string][]string{
			//"WedgeLeaf" : strings.Split("p p A0.6 y Colour255,255,255 . [ F Y . F A1.5708 Y Y A0.6 y Colour255,255,255 . . F Y Colour255,255,255 Colour255,200,200 . F ] Colour255,255,255 . A0.523598", " "),
			"WedgeLeaf": strings.Split("p p A0.6 y Colour255,255,255 . [ F Y . F A1.5708 Y Y A0.6 y Colour255,255,255 . . F Y Colour255,255,255 Colour255,200,200 . F ] Colour255,255,255 .", " "),
		},
		"Square1": map[string][]string{
			//"chunk" : strings.Split("Colour255,0,0 A1.5707 [ .  F Y . F Y . . F Y . F Y . ]", " "),
			"chunk1":  strings.Split("A1.5707 Colour255,255,255 [ [ P s s F . ] F y Colour0,255,0 . F y . . F y . F y [ P s s F Colour255,255,255 . ] ]", " "),
			"chunk":   strings.Split("A1.5707 Colour255,255,255 s s s s [ chunk1 ] Y  Colour255,255,255 [ chunk1 ] Y [ chunk1 ] Y [ chunk1 ]", " "),
			"Square1": strings.Split("[ chunk ]", " "),
		},
		"Face": map[string][]string{
			"chunk": strings.Split("Colour255,0,0 A1.5707 [ . [ F . ] Y [ F  . ] ]", " "),
			"Face":  strings.Split("s s s s s s chunk Y chunk Y chunk Y chunk", " "),
		},
		"Arrow": map[string][]string{
			"chunk": strings.Split("Colour255,0,0 A1.5707 [ [ F op [ F Colour0,255,0 op ] ] Y [ Colour0,0,255 F  op ] ]", " "),
			"Arrow": strings.Split("s s s s chunk R chunk R chunk R chunk", " "),
		},
		"Tetrahedron": map[string][]string{
			"c1": s("[ F P F Y F . ]"),
			"c2": s("[ f P f Y F . ]"),
			"c3": s("[ f P F Y f . ]"),
			"c4": s("[ F P f Y f . ]"),
			"Tetrahedron": s(`A1.5707 Colour255,0,0
            [ c1 c2 c3 ] Colour0,255,0
            [ c1 c3 c4 ] Colour0,0,255
            [ c1 c2 c4 ] Colour255,255,255
            [ c4 c3 c2 ]
        `),
		},
		"Prism": map[string][]string{
			"c1": s("[ . ]"),
			"c2": s("[ F . ]"),
			"c3": s("[ P F . ]"),
			"c4": s("[ Y F y . ]"),
			"c5": s("[ Y F y F . ]"),
			"c6": s("[ Y F y P F . ]"),
			"Prism": s(`Colour0,255,0  A1.5707 hs y F Y p F P hS
                [ c1 c2 c3 ] [ c1 c4 c2 ] [ c1 c3 c4 ]
                [ c4 c6 c5 ] [ c4 c5 c2 ] [ c4 c3 c6 ]
                [ c2 c5 c6 ] [ c6 c3 c2 ]
                hs F Y F y hS p`),
		},
		"Prism1": map[string][]string{
			"c1": s("[ . ]"),
			"c2": s("[ F . ]"),
			"c3": s("[ P F . ]"),
			"c4": s("[ Y F y . ]"),
			"c5": s("[ Y F y F . ]"),
			"c6": s("[ Y F y P F . ]"),
			"Prism1": s(`Colour255,0,0 A1.5707 hs y F Y p F P hS
                [ c1 c2 c3 ] [ c1 c4 c2 ] [ c1 c3 c4 ]
                [ c4 c6 c5 ] [ c4 c5 c2 ] [ c4 c3 c6 ]
                [ c2 c5 c6 ] [ c6 c3 c2 ]
                hs F Y F y hS p`),
		},
		"Circle": map[string][]string{
			"spoke":  s("spoke Y [ F F s s s Tetrahedron ]"),
			"Circle": s(`A0.71 spoke`),
		},
		"Icosahedron": map[string][]string{
			"patch1": s(`
                        [ F F Y F . ] [ P F F p F . ] [ Y F F P F . ] reverseTriangle
                        [ F F y F . ] [ P F F p F . ] [ y F F P F . ]
                    `),
			"ring": s(`
                Colour255,0,0 [ F [ F [ Y F . ] [ y F . ] ] [ P F F . ] ] reverseTriangle
                Colour0,255,0 [ F [ F [ Y F . ] [ y F . ] ] [ p F F . ] ]
                Y Y
                Colour0,0,255 [ F [ F [ Y F . ] [ y F . ] ] [ P F F . ] ] reverseTriangle
                Colour255,0,255 [ F [ F [ Y F . ] [ y F . ] ] [ p F F . ] ]
`),
			"Icosahedron": s(`A1.5707
                    [ Y Y p R R p
                    Colour255,255,255
                    patch1
                ]
                [ ring ]
                [
                    Y P
                    ring
                ]
                [
                    Y R
                    ring
                ]
                [
                    Colour255,255,0
                    patch1
                    R R
                    Colour255,0,255
                    patch1
                    P P
                    Colour0,255,255
                    patch1
                ]
                `),
		},
	}
	return rb

}

func rules() map[string][]string {
	rules := map[string][]string{
		"st": strings.Split("R Y T F", " "),
		"le": strings.Split("[ y y P P Colour0,255,0 T y R R R Colour0,255,0 T ]", " "),

		"Koch10":  strings.Split("A1.57 Koch10l", " "),
		"Koch10l": strings.Split("Koch10l Y Koch10r Y", " "),
		"Koch10r": strings.Split("y Koch10l y Koch10r", " "),
		"stem":    []string{"F", "A0.35", "R", "[", "s", "square", "stem", "]", "[", "A0.75", "y", "s", "square", "stem", "]"},
	}
	return rules
}

func runRules(commands []string, rules map[string][]string, recursion int) []string {

	//if (compiled ==nil) {
	//fmt.Println("Compiling")
	for i := 0; i < recursion; i++ {
		//fmt.Println("Recursing")
		i = i + 0 //Must use the loop variable or the go compiler complains
		commands = rewrite(commands, rules)
	}
	compiled = commands
	//} else {
	////fmt.Println("Using compiled")
	//commands =  compiled
	//clrClr = 0
	//}
	return commands
}

func runRuleset(aString []string, ruleset map[string]map[string][]string) []string {
	var outString []string
	for i := 0; i < 50; i++ {
		outString = []string{}
		for _, orig := range aString {
			new, ok := ruleset[orig]
			if ok {
				//fmt.Println("Replace:")
				//fmt.Println(orig, new)
				outString = append(outString, runRules([]string{orig}, new, ruleBookMeta()[orig].iterations)...)
			} else {
				//fmt.Println(orig, new)
				outString = append(outString, []string{orig}...)
			}
		}
		//fmt.Println("Iterate:")
		//fmt.Println(outString)
		aString = outString
	}
	return outString
}

func Expand(seed string) []string {
	start := s(seed)
	out := runRuleset(start, ruleBook())
	return out
}

func ruleBookMeta() map[string]ruleMeta {
	ret := map[string]ruleMeta{
		"Quad":        ruleMeta{iterations: 3, x: 0.5, y: 0.5, style: TRIANGLES},
		"risingVine":  ruleMeta{iterations: 10, x: 0.5, y: 0.5, style: TRIANGLES},
		"Koch3":       ruleMeta{iterations: 4, x: 0.5, y: 0.5, style: LINES},
		"Koch2":       ruleMeta{iterations: 4, x: 0.5, y: 0.5, style: TRIANGLES},
		"3DTreeLeafy": ruleMeta{iterations: 3, x: 0.5, y: 0.9, style: TRIANGLES},
		"3DTree3":     ruleMeta{iterations: 4, x: 0.5, y: 0.9, style: TRIANGLES},
		"Tree3":       ruleMeta{iterations: 4, x: 0.5, y: 0.9, style: LINE_LOOP},
		"starburst":   ruleMeta{iterations: 10, x: 0.5, y: 0.7, style: TRIANGLES},
		"leaf2":       ruleMeta{iterations: 7, x: 0.5, y: 0.7, style: TRIANGLE_FAN},
		"leaf":        ruleMeta{iterations: 7, x: 0.5, y: 0.7, style: TRIANGLES},
		"ilake":       ruleMeta{iterations: 4, x: 0.5, y: 0.7, style: TRIANGLES},
		"lineStar":    ruleMeta{iterations: 8, x: 0.5, y: 0.7, style: LINE_STRIP},
		"KIomega":     ruleMeta{iterations: 3, x: 0.5, y: 0.5, style: TRIANGLES},
		"Gosper":      ruleMeta{iterations: 6, x: 0.5, y: 0.5, style: LINE_LOOP},
		"Sierpinksi":  ruleMeta{iterations: 7, x: 0.5, y: 0.5, style: LINE_LOOP},
		"Square":      ruleMeta{iterations: 1, x: 0.5, y: 0.5, style: TRIANGLES},
		"Square1":     ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"Flower":      ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"Flower12":    ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"Flower11":    ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"Flower10":    ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"Plant":       ruleMeta{iterations: 5, x: 0.5, y: 0.5, style: TRIANGLES},
		"FlowerField": ruleMeta{iterations: 30, x: 0.5, y: 0.5, style: TRIANGLES},
		"WedgeLeaf":   ruleMeta{iterations: 1, x: 0.5, y: 0.5, style: LINE_LOOP},
		"Face":        ruleMeta{iterations: 4, x: 0.5, y: 0.5, style: TRIANGLES},
		"Arrow":       ruleMeta{iterations: 4, x: 0.5, y: 0.5, style: TRIANGLES},
		"Tetrahedron": ruleMeta{iterations: 2, x: 0.5, y: 0.5, style: TRIANGLES},
		"Circle":      ruleMeta{iterations: 10, x: 0.5, y: 0.5, style: TRIANGLES},
		"Prism":       ruleMeta{iterations: 3, x: 0.5, y: 0.5, style: TRIANGLES, backfaceCull: true, useLighting: true},
		"Prism1":      ruleMeta{iterations: 3, x: 0.5, y: 0.5, style: TRIANGLES, backfaceCull: true, useLighting: true},
		"Icosahedron": ruleMeta{iterations: 4, x: 0.5, y: 0.5, style: TRIANGLES, backfaceCull: true, useLighting: true},
	}
	return ret
}

func tr(trans mgl32.Mat4) []float32 {
	polyCount = polyCount + 1
	vec := mgl32.Vec4{0.0, 1.0, 0.0, 1}
	p1 := trans.Mul4x1(vec)
	vec = mgl32.Vec4{-0.1, 0.0, 0.0, 1}
	p2 := trans.Mul4x1(vec)
	vec = mgl32.Vec4{0.1, 0.0, 0.0, 1}
	p3 := trans.Mul4x1(vec)
	newTri := []float32{p1[0], p1[1], p1[2], p2[0], p2[1], p2[2], p3[0], p3[1], p3[2]}
	return newTri
}

func line(trans mgl32.Mat4) []float32 {
	polyCount = polyCount + 1
	vec := mgl32.Vec4{0.0, 1.0, 0.0, 1}
	p1 := trans.Mul4x1(vec)
	vec = mgl32.Vec4{-0.1, 0.0, 0.0, 1}
	p2 := trans.Mul4x1(vec)
	newTri := []float32{p1[0], p1[1], p1[2], p2[0], p2[1], p2[2]}
	return newTri
}

func point(trans mgl32.Mat4) []float32 {
	polyCount = polyCount + 1
	vec := mgl32.Vec4{0.0, 0.0, 0.0, 1}
	p1 := trans.Mul4x1(vec)
	newVert := []float32{p1[0], p1[1], p1[2]}
	return newVert
}

func S(str string) []string {
	return s(str)
}
