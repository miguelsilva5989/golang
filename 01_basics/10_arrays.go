package main

import (
	"fmt"
)

func main() {
	// arrays max size is defined when initialized
	var arr [5]string
	fmt.Println(arr)

	var arr2 [5]int
	fmt.Println(arr2)

	// print pos 0 of arr2
	arr2[0] = 5
	fmt.Println(arr2[0])
	fmt.Println(arr2)

	//define elements when initalizing
	arr3 := [3]int{4, 5, 6}
	fmt.Println(arr3)

	fmt.Println("length", len(arr3))

	sum := 0
	for _, i := range arr3 {
		sum += i
	}
	fmt.Println("sum of elements", sum)

	//2d array
	arr2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d)

}
