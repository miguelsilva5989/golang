package main

import (
	"fmt"
)

func main() {
	// Slices
	// pointer - which position to slice from
	// length - length of slice 
	// capacity - how many elements the slice would be able to hold
	// starting from the pointer until its end
	// i.e: array of 4 elements, pointer in element 1, length 2
	// capacity would be 3

	var arr [5]int = [5]int{1,2,3,4,5}
	var slice []int = arr[1:3]

	fmt.Println(slice)
	fmt.Println(arr[:]) //slice on the go
	fmt.Println(arr[1:2])

	fmt.Println("len", len(slice))
	fmt.Println("capacity", cap(slice))

	// extend slice to be full length and adds other elements
	fmt.Println(slice[:cap(slice)])
	fmt.Println(slice)
	// slice the slice
	fmt.Println(slice[1:])

	// create slice without using an array
	var a[]int = []int{5,6,7,8,9}
	fmt.Println("capacity", cap(a))
	fmt.Println(cap(a[:3]))

	// increase slice size (actually it creates a new slice with more values)
	b := append(a, 10)
	fmt.Printf("type %T\n", b)
	fmt.Println(b)
	// or the same slice var as before
	a = append(a, 10)
	fmt.Println(a)

	//using make to create a slice
	c := make([]int, 5)
	fmt.Println(c)
	
}