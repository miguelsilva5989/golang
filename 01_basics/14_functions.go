package main

import (
	"fmt"
)

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

// return using var names
func addSub2(x, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

// return using var names
func addSub3(x, y int) (z1 int, z2 int) {
	defer fmt.Println("example on doing action before return using defer")
	z1 = x + y
	z2 = x - y
	fmt.Println("I'm printed before the defer statement")
	return
}

func main() {
	add, sub := addSub(2, 3)
	fmt.Println(add, sub)

	add2, sub2 := addSub2(2, 3)
	fmt.Println(add2, sub2)

	addSub3(2, 3)
}
