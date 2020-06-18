package main

import (
	"fmt"
)

func test(x int) {
	fmt.Println("test", x)
}

// function which receives function
func funcTest(myFunc func(int) int) {
	fmt.Println("test2", myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	a := test
	a(2)

	test2 := func(x int) int {
		return x * -1
	}
	fmt.Println(test2(2))

	test3 := func(x int) int {
		return x * -1
	}(2) //notice the int here which calls the function in-line
	fmt.Println(test3)

	// send function to funcTest
	funcTest(test2)

	// function closure
	returnFunc("returned a function and executed it in-place")()
}
