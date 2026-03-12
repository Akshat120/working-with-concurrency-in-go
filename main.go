package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

// main fun itself a goroutine
// goroutines are light-weight threads,
// group of goroutine managed by go-scheduler
func main() {
	// below go routine will run in its own go-routine
	// To see the output of the below goroutine, main goroutine needs
	// to be waited for sometime.
	go printSomething("This is first thing to be printed!")

	// Bad Solution
	time.Sleep(1 * time.Second)

	// this will be run in main go-routine
	printSomething("This is second thing to be printed!")

}
