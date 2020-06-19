package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	area() float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	fmt.Println(c1.area())

	r1 := Rectangle{2, 4}
	fmt.Println(r1.area())

	// shapes := []type{c1, r1} // what type to use here?
	// that's where interface comes in

	shapes := []Shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("type: %T\n", shape)
		fmt.Println("shape area: ", shape.area())
		fmt.Println("shape area using func: ", getArea(shape))
	}
}
