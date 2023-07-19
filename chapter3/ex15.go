package main

import "fmt"

func main() {
	// var dataStream chan interface{}  --> nil chan

	var receiveChan <-chan interface{}
	var sendChan chan<- interface{}
	dataStream := make(chan interface{}) // non nil chan

	receiveChan = dataStream
	sendChan = dataStream

	go func() {
		sendChan <- "hello chan"
	}()

	fmt.Println(<-receiveChan)
}
