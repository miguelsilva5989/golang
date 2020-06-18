package main

import "fmt"

func changeValue(str *string) {
	// * before parameter - means it receives the pointer to the var
	// this way we can change the variable using the pointer
	*str = "changed!"
}

func main() {
	// & get the pointer
	// * dereference pointer

	x := 7
	fmt.Println(&x) // where is value 7 stored in memory

	y := &x           // y is pointing to the reference of x in memory
	fmt.Println(x, y) // where is value 7 stored in memory

	*y = 8            // dereferenced the pointer, and change that pointer to another value
	fmt.Println(x, y) // and now the value of x will be 8

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	var pointer *string = &toChange
	fmt.Println(pointer)
	fmt.Println(*pointer)          // value at the pointer
	fmt.Println(pointer, &pointer) // pointer and the pointer to the pointer variable
}
