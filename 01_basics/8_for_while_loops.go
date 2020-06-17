package main

import (
	"fmt"
)

func main() {
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	for true {
		if count == 5 {
			fmt.Println("reached number 5")
			break
		}
		count++
	}

	for count <= 10 {
		fmt.Println(count)
		count++
	}

	for i, ch := range "abc" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	for i := 0; i < 3; i++ {

		if i%2 == 0 {
			fmt.Println(i, "is divisible by 2")
			continue
		}

		fmt.Println(i, "is not divisible")
	}
}
