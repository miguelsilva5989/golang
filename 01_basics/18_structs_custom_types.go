package main

import (
	"fmt"
)

type Point struct {
	x int32
	y int32
}

// Circle embedded Struct
type Circle struct {
	radius float64
	center *Point
}

type Circle2 struct {
	radius float64
	*Point // can also be declared without a variable name (i.e.:e center)
	// as long as the variables x,y do not exist in Circle2
}

func changeX(pt *Point) {
	pt.x = 100 // structs don't need to be dereferenced to change values
	// i.e.: a normal variable would be changed as *str = "changed!" (check 17_pointer_dereference.go)
}

func main() {
	var p1 Point = Point{1, 2}
	fmt.Println(p1.x, p1.y)
	p1.x = 10 // change value of x
	fmt.Println(p1.x)

	p2 := Point{x: 5} // no y defined
	fmt.Println(p2)   // y has default int32 value which is 0

	p3 := &Point{y: 2} // pointer to the values of point
	fmt.Println(p3)
	changeX(p3)
	fmt.Println(p3)

	c1 := Circle{3.75, &Point{4, 5}}
	fmt.Println(c1)
	fmt.Println(c1.center)
	fmt.Println(c1.center.x)

	c2 := Circle2{2.75, &Point{1, 2}}
	fmt.Println(c2.x) // Circle2 does not have center
}
