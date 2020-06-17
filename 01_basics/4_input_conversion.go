package main

import (
	"bufio" //input
	"fmt"
	"os"
	"strconv"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin) // user input
	// fmt.Print("Type something: ")
	// scanner.Scan()
	// input := scanner.Text()

	// fmt.Printf("You typed: %q", input)

	scanner := bufio.NewScanner(os.Stdin) // user input
	fmt.Print("Type the year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32) //convert string to int, _ is required in case of error

	fmt.Printf("You will be %d years old at the end of 2020", 2020-input)
}
