package main

import (
	"fmt"

	"github.com/miguelsilva5989/port_scanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	// open := port.ScanPort("tcp", "localhost", 1313)
	// fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
