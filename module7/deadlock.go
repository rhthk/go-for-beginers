package module7

import "fmt"

func deadlock() {
	// unbuffered channel
	ch := make(chan int)
	go func() {
		// add receiver before adding values to channel
		fmt.Println(<-ch)
	}()
	// add values to channel
	ch <- 1
}
