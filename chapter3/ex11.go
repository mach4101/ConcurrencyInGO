package main

import "sync"

var (
	onceA, onceB sync.Once
	initB        func()
)

func main() {
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }
	onceA.Do(initA)
}
