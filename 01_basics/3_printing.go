package main

import "fmt"

func main() {
	fmt.Printf("Type %T", 10)
	fmt.Printf("\nType %T value is %v", 10, 10)

	var x string = fmt.Sprintf("\nType %T value is %v", 10, 10)
	fmt.Println("\nSprint (stored in var)", x)

	fmt.Printf("num2: %f", 2.37452435123123)
	fmt.Printf("\nnum3: %g", 2.37452435123123)

	fmt.Printf("\nString: %s", "Zeca")
	fmt.Printf("\nString: %q", "Zeca")

	// Padding and width
	fmt.Printf("\n%9s", "Mike")          // with of 9 with lpadding
	fmt.Printf("\n%-9s is cool", "Mike") // right justified
	fmt.Printf("\n%.2f", 2.465465)       // precision with of 2
	fmt.Printf("\n%07d", 45)             // left padding of 0
}
