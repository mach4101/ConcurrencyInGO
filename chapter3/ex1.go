package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(time.Duration(1 * time.Millisecond))
}

func sayHello() {
	fmt.Println("hello")
}
