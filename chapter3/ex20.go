package main

import "fmt"

func main() {
	resultStream := chanOwner()

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}

func chanOwner() <-chan int {
	resultStream := make(chan int, 20)

	go func() {
		defer close(resultStream)

		for i := 0; i <= 20; i++ {
			resultStream <- i
		}
	}()
	return resultStream
}
