package main

import (
	"fmt"
	"sync"
)

var (
	count      int
	once       sync.Once
	increments sync.WaitGroup
)

func main() {
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("count is %d\n", count)
}

func increment() {
	count++
}
