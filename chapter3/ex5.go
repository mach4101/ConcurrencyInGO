package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
}

func main() {
	const numGreets = 5
	var wg sync.WaitGroup

	wg.Add(numGreets)

	for i := 0; i < numGreets; i++ {
		go hello(&wg, i+1)
	}

	wg.Wait()
}
