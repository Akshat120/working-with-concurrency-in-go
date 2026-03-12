package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

// main fun itself a goroutine
// goroutines are light-weight threads,
// group of goroutine managed by go-scheduler
func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	// Increased waitgroup counter to the len of words
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%v: %v", i, x), &wg)
	}

	// Bad Solution
	// if the slice of words we have created above,
	// if its size increases then we don't know how
	// much we need to wait for all the goroutine
	// needs to be completed going to variable time.
	// time.Sleep(1 * time.Second)

	// Waiting on waitgroups
	wg.Wait()

	wg.Add(1)
	// this will be run in main go-routine
	printSomething("This is second thing to be printed!", &wg)

}
