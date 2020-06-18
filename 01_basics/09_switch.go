package main

import (
	"fmt"
)

func main() {
	ans := 10

	switch ans {
	case 10:
		fmt.Println("is 10")
	default:
		fmt.Println("nope")
	}

	switch {
	case ans < 10:
		fmt.Println("less than 10")
	default:
		fmt.Println("nope")
	}
}
