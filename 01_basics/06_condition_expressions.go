package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6.5
	z := 5
	
	val := x <= 4
	fmt.Printf("%t", val)

	val = float64(x) >= y
	fmt.Printf("\n%t", val)

	val = x == z
	fmt.Printf("\n%t", val)

	val = x != z
	fmt.Printf("\n%t", val)

	fmt.Printf("\n%t", x == z)

	a := "ze"
	b := "Ze"
	c := "zee"

	val = a <= b
	fmt.Printf("\n%t", val)

	val = a <= c
	fmt.Printf("\n%t", val)

	fmt.Printf("\n%t", a < b && a < c)
	fmt.Printf("\n%t", a < b || a < c)
	fmt.Printf("\n%t", !true == false)
}