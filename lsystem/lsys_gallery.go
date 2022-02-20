// lsys_gallery.go
package lsystem

var renderGallery []string = []string{
	//`[ mirrorOff noRtt Colour255,0,0 s s A0.78 R A1.57 Arrow P Arrow F F s s F F S S p A0.78 p A0.01 Y F F Texture(0) ] TextureOff`,
	`[ mirrorOff noRtt Colour255,0,0 Texture(0) ] `,
}

func InitGallery() []string {

	gallery := []string{
		//`[ mirrorOff noRtt Colour255,0,0 s s A0.78 R A1.57 Arrow P Arrow F F s s F F S S p A0.78 p A0.01 Y F F Texture(0) ] TextureOff`,
		`[ mirrorOff noRtt Colour255,0,0 Texture(0) ] `,
		`noRtt [ s s f S S s s s s A1.5707 R R Y F y f f s s s HR HR [ [ P Circle ] Plant ] ]`,
		"A1.57 s s f S S S HR R s s s s s s s [ Colour0,255,0 3DTree3 ] A1.57 [ Y F F F F F F F F F F y s r Colour255,0,0  3DTreeLeafy ] A1.57 [ y F F F F F F F F F F Y s r r Colour0,0,255 3DTreeLeafy ] Colour0,255,255 A1.57  [ S P [ leaf2 ] A1.57 Y Y [ leaf2 ] ] ",
		`[ mirrorOff noRtt LightsOn Colour255,0,0 f A0.35 R HR [ s s
            Arrow F Arrow Prism [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
            A1.5707
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(0) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(1) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(2) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(3) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(4) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(5) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(6) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(7) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(8) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(9) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(10) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(11) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(12) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(13) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(14) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(15) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(16) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(17) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(18) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(19) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(20) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(21) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(22) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
    ]
]

`,
		/*`[ mirrorOff noRtt Colour255,0,0 HR HR HR HR Hinge(0) s s s A0.7 [ P F F Quad ] ]`,
		//[ mirrorOn Scale(-1.0,1.0,1.0) s s s s s A1.5707 R R Y F y f f s s s HR HR [ [ P Circle ] Plant ] mirrorOff ]`,
		"A1.5707 Colour255,0,0 P P Square1",

		`s s s s [ [                 [ A1.5707 Colour255,0,0 [ Tetrahedron ] P [ P P HR Circle ] [ P HR Circle ] HR Tetrahedron ] ] ]`,
		//[ mirrorOn Scale(1.0,-1.0,1.0) A1.5707 Colour255,0,0 LightsOn [ Tetrahedron ] P [ P P Circle ] [ P Circle ] Tetrahedron mirrorOff ] ] ]`,

		`s s s s HR HR HR HR [ s F F                            A1.5707 Colour255,0,0 LightsOn Icosahedron LightsOff ]`,

		"A1.5707 HR S [ Arrow Y P ] Y Y Arrow",
		"A1.5707 Colour255,0,0 s s s s s s HR HR HR ilake",
		"s HR s s s F FlowerField",
		"s s s s s s s starburst",
		"A1.5707 Colour255,0,0 s s s s s s HR orient",
		"A1.5707 Colour255,0,0 HR P HR Face",
		"A1.5707 Colour255,0,0 HR P HR Arrow",
		"A1.5707 s s R R P A0.7 P Arrow",
		"A1.5707 s s s s R R A0.785 P Flower",
		"A1.5707 s s s s s [ R R s Flower11 ] [ Y F F y R R s Flower11 ] [ y F F Y R R s Flower10 ]",
		"A1.57 R [ s s F F F F F F F F P s s Colour200,200,200 ] H R [ s Colour0,0,200 [ S S S square starburst ] ] s s s s s [ Colour0,255,0 ] A1.57 [ Y F F F F F F F F F F y s r Colour255,0,0  ] A1.57 [ y F F F F F F F F F F Y s r r Colour0,0,255 ] Colour0,255,255 A1.57  [ S P [ leaf2 ] A1.57 Y Y [ leaf2 ] ] ",
		"A1.57 R R HR R s [ s Colour0,0,50 starburst ] s s s s s  A1.57  A1.57  Colour0,255,255 A1.57  [ S P  A1.57 Y Y [ leaf2 ] ] ",
		"s s s A1.50 Colour255,0,0 starburst Y Y starburst P F F F Colour0,255,0 starburst Y Y starburst",
		"HR s s s s s s 3DTreeLeafy",
		"s s s s s s s Tree3",
		"A1.57 s s s s lineStar",
		"A1.57 s s s s leaf",
		"A1.57 s s s s leaf2",
		"A0.7 R Y Y s s s s s s s s s s Koch3",
		"A0.7 R P Y s s s s s s s Gosper",
		"A0.7 R P s s s s s s s s Sierpinksi",
		"s s s s s s s s Koch2",
		"s s s KIomega",*/
	}
	return gallery

}

var PlantGallery []string = []string{
	`s s s s [ [                 [ A1.5707 Colour255,0,0 [ Tetrahedron ] P [ P P HR Circle ] [ P HR Circle ] HR Tetrahedron ] ] ]`,
	`s s s s HR HR HR HR [ s F F                            A1.5707 Colour255,0,0 LightsOn Icosahedron LightsOff ]`,
	"A1.5707 HR S [ Arrow Y P ] Y Y Arrow",
	"A1.5707 Colour255,0,0 s s s s s s HR HR HR ilake",
	"s HR s s s F FlowerField",
	"s s s s s s s starburst",
	"A1.5707 Colour255,0,0 s s s s s s HR orient",
	"A1.5707 Colour255,0,0 HR P HR Face",
	"A1.5707 Colour255,0,0 HR P HR Arrow",
	"A1.5707 s s R R P A0.7 P Arrow",
	"A1.5707 s s s s R R A0.785 P Flower",
	"A1.5707 s s s s s [ R R s Flower11 ] [ Y F F y R R s Flower11 ] [ y F F Y R R s Flower10 ]",
	"A1.57 R [ s s F F F F F F F F P s s Colour200,200,200 ] H R [ s Colour0,0,200 [ S S S square starburst ] ] s s s s s [ Colour0,255,0 ] A1.57 [ Y F F F F F F F F F F y s r Colour255,0,0  ] A1.57 [ y F F F F F F F F F F Y s r r Colour0,0,255 ] Colour0,255,255 A1.57  [ S P [ leaf2 ] A1.57 Y Y [ leaf2 ] ] ",
	"A1.57 R R HR R s [ s Colour0,0,50 starburst ] s s s s s  A1.57  A1.57  Colour0,255,255 A1.57  [ S P  A1.57 Y Y [ leaf2 ] ] ",
	"s s s A1.50 Colour255,0,0 starburst Y Y starburst P F F F Colour0,255,0 starburst Y Y starburst",
	"HR s s s s s s 3DTreeLeafy",
	"s s s s s s s Tree3",
	"A1.57 s s s s lineStar",
	"A1.57 s s s s leaf",
	"A1.57 s s s s leaf2",
	"A0.7 R Y Y s s s s s s s s s s Koch3",
	"A0.7 R P Y s s s s s s s Gosper",
	"A0.7 R P s s s s s s s s Sierpinksi",
	"s s s s s s s s Koch2",
	"s s s KIomega"}

var snekGallery []string = []string{
	`[ mirrorOff noRtt LightsOn Colour255,0,0 f A0.35 R HR [ s s
            Arrow F Arrow Prism [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
            A1.5707
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(0) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(1) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(2) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(3) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(4) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(5) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(6) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(7) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(8) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(9) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(10) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(11) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(12) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(13) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(14) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(15) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(16) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(17) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(18) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(19) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(20) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(21) Prism  [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
                hs F Scale(1.2,1.2,1.2) p F Scale(0.8333,0.8333,0.8333) hS Hinge(22) Prism1 [ hs F  A2.356 P  Scale(1.0,1.2,1.0) Qua ]
    ]
]

`,
}
