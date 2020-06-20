// https://www.youtube.com/watch?v=H85TJW_4JgM

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup with error:", r)
	}
}

func say(s string) {
	defer wg.Done() // tell waitgroup that it's finished
	defer cleanup() // clean up runs before wg.Done() -> LIFO
	// wg.Done() could also be done inside the cleanup

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("oh my god! not number 2!!!")
		}
	}
}

func main() {
	wg.Add(1)
	go say("hey") // goroutine (lightweight thread)
	wg.Add(1)
	go say("there")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
