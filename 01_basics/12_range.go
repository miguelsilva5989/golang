package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4, 6, 8, 9, 12, 2}

	for i, elem := range a {
		fmt.Printf("%d: %d\n", i, elem)
	}

	// if we don't need the index, use the _ so no error occurs
	for _, elem := range a {
		fmt.Printf("%d ", elem)
	}

	var b []int = []int{1, 2, 4, 2, 2, 5}
	// (naive approach) print the duplicates in slice using range
	for i, elem := range b {
		for j, elem2 := range b {
			if elem == elem2 && i > j {
				fmt.Println(elem)
			}
		}
	}

	// (better approach) print the duplicates in slice using range
	for i, elem := range b {
		for j := i + 1; j < len(b); j++ {
			elem2 := b[j]
			if elem == elem2 {
				fmt.Println(elem)
			}
		}
	}
}
