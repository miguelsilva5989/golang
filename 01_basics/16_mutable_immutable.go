package main

import (
	"fmt"
)

func changeFirst(slice []int) {
	slice[0] = 100
}

func main() {
	// x is immutable
	var x int = 5
	y := x
	y = 7
	fmt.Println(x, y)

	// arrays are immutable
	var d [2]int = [2]int{3, 4}
	e := d
	e[0] = 100
	fmt.Println(d, e)

	// z (slice) is mutable
	var z []int = []int{3, 4, 5} // pointer to slice
	w := z                       // actually pointing to the slice
	w[0] = 100
	fmt.Println(z, w, "z was also changed as a slice is mutable")

	// map is also mutable
	var a map[string]int = map[string]int{
		"hello": 3,
	}
	b := a
	b["y"] = 2
	fmt.Println(a, b)

	// change slice inside function
	var g []int = []int{1, 2, 3}
	fmt.Println(g)
	changeFirst(g)
	fmt.Println(g)

}
