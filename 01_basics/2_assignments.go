package main

import "fmt"

func main() {
	// implicit define
	var num1 = 250

	// explicit
	var num2 uint16 = 6999

	// expression assignment operator aka walrus operator
	num3 := 6
	// num3 := "nope" - would return error as num3 was a int before

	fmt.Printf("guess type %T", num1)
	fmt.Printf("\ntype %T", num2)
	fmt.Printf("\nguess type %T\n", num3)

	//Booleans

	ze := false
	fmt.Println(ze)

	var bl bool
	fmt.Println("default bool", bl)

	// no assignment

	var empty uint64
	fmt.Println("default uint", empty)

}
