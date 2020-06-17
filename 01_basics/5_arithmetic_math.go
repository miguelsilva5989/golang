package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 3
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Println(answer)

	var num3 int = 9
	var num4 int = 4
	answer2 := num3 / num4 // answer will be an int as the vars are int
	fmt.Printf("%d", answer2)

	var num5 float32 = 9
	var num6 float32 = 4
	answer3 := num5 / num6 // answer will be an float32 as the vars are float32

	fmt.Printf("\n%g", answer3)

	var num7 int = 9
	var num8 int = 4
	answer4 := num7 % num8
	fmt.Printf("\nremainder %d\n", answer4)

	fmt.Println(math.Exp(4))
	fmt.Println(math.Pow(2, 10))
}