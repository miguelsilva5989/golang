// https://www.youtube.com/watch?v=S11VFAMEs6E

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5 //send value to channel
}

func main() {
	// create channel
	fooVal := make(chan int, 10) // buffer for 10 items

	// go foo(fooVal, 5)
	// go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal
	// fmt.Println(v1)
	// fmt.Println(v2)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	wg.Wait()
	// to prevent deadlock on channel - we need to close channel
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

}
