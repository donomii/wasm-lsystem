package lsystem

import (
	"time"

	"../tween"
)

//Init global data for the snek display
var (
	hinges    []float32
	frog      []float32
	tapering  []float32
	snowflake []float32
	terrier   []float32
	rooster   []float32
	goldfish  []float32
	snek      []float32
	ball      []float32
	bow       []float32
)

func init_snek() {
	hinges = []float32{2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5, 2 * 3.1415927 * 0.5}
	//start_clock(&clock, 0.0, 1.0, 10.0)
	//start_linear_animation(&hinges[1], 0.0, 10.0, 100.0)
	frog = []float32{-1, 0, -1, 1, 1, 2, -1, 1, 0, 0, 1, 2, -1, 0, 0, -1, 1, 2, -1, -1, 1, 0, 1}
	tapering = []float32{0, 1, 0, 2, 2, 0, -1, 2, -1, -1, 0, 0, 0, 1, -1, 2, -1, 0, 2, 2, 0, -1, 0}
	snowflake = []float32{-1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1}
	rooster = []float32{2, 0, -1, 0, -1, 1, 2, 1, 0, 2, 2, 0, 1, 2, 1, -1, 0, -1, 0, 2, 2, 0, 0}
	terrier = []float32{2, 0, 2, 2, 0, 2, 0, 0, 0, 2, 2, 0, 2, 0, 0, 2, 0, 2, 2, 0, 0, 0, 0}
	ball = []float32{-1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, -1, -1}
	snek = []float32{0, 1, 0, 2, 2, 0, 2, 2, 0, 1, 0, 0, 0, 2, -1, -1, 2, -1, 0, 0, 2, 1, -1}
	goldfish = []float32{0, 0, 2, 0, 2, 0, 0, 2, -1, 1, 1, 2, -1, 1, 0, 1, -1, 2, 1, 1, -1, 0, 2}
	bow = []float32{-1, -1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, 1}
	for i, v := range snek {
		hinges[i] = (v + 2) * 3.141527 / 2.0 //+2 because the prism primitive leaves the rotation "upside down"
	}

}

func transform_snek(s *Scene) {
	rubiks := [][]float32{tapering, snowflake, rooster, terrier, frog, bow, goldfish, ball}
	for _, critter := range rubiks {
		if s.Active {
			tween.StartLinear(&s.Clock, 0.0, 1.0, 25.0)
			for i, v := range critter {
				if hinges[i] != (v+2)*3.141527/2.0 {
					time.Sleep(2 * time.Second)
					tween.StartLinear(&hinges[i], hinges[i], (v+2)*3.141527/2.0, 1.0)
				}
				//hinges[i] = (v+2)*3.141527/2.0  //+2 because the prism primitive leaves the rotation "upside down"
			}
			tween.StartLinear(&s.Clock, 0.0, 1.0, 10.0)
		}
		time.Sleep(10.0 * time.Second)
	}

	transform_snek(s)
}
