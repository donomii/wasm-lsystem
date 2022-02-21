package lsystem

func Icosahedron() string {
	return `Colour254,254,254 deg30 Arrow  F 
	s s s s
	[ s s HR Icosahedron ] 
	`
}
func AllObjects() string {
	return `Colour254,254,254 deg30 r r F p p p f f f f [
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

`
}
