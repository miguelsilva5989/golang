package main

import (
	"fmt"
)

func main() {
	var mp map[string]int = map[string]int{
		"ze":  1,
		"ze2": 5,
		"ze3": 7,
	}
	fmt.Println(mp)

	mp2 := map[string]int{
		"ze":  1,
		"ze2": 5,
		"ze3": 7,
	}
	fmt.Println(mp2)

	mp3 := make(map[string]int)
	mp3 = map[string]int{
		"ze":  1,
		"ze2": 5,
		"ze3": 7,
	}
	fmt.Println(mp3)

	// access and change values
	fmt.Println("ze:", mp["ze"])
	mp["ze"] = 100
	fmt.Println("ze is now:", mp["ze"])

	// delete
	delete(mp, "ze")
	fmt.Println(mp)

	// check if key exists
	val, ok := mp["ze2"]
	fmt.Println("value of ze2:", val, "exists:", ok)

	fmt.Println("len:", len(mp))

}
