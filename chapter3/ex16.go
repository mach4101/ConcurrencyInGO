package main

import "fmt"

func main() {
	valueStream := make(chan interface{})
	close(valueStream)

	intStream := make(chan int)
	close(intStream)

	integer, ok := <-intStream
	fmt.Printf("(%v) : %v\n", ok, integer)
}
