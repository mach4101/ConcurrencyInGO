package main

import (
	"fmt"
	"time"
)

// range for nil channel, causing memory leakage

///func main() {
///	var wg sync.WaitGroup
///	wg.Add(1)
///	doWork := func(strings <-chan string, wg *sync.WaitGroup) <-chan interface{} {
///		completed := make(chan interface{})
///
///		go func() {
///			defer wg.Done()
///			defer fmt.Println("doWork exited!")
///			defer close(completed)
///			for s := range strings {
///				fmt.Println(s)
///			}
///		}()
///		return completed
///	}
///
///	doWork(nil, &wg)
///	fmt.Println("Done.")
///	wg.Wait()
///	fmt.Println("program exit!")
///}

func main() {
	doWork := func(done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exit!")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("canceling dowork goroutine")
		close(done)
	}()

	<-terminated
	fmt.Println("done!")
}
