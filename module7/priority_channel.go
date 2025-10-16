package module7

import (
	"fmt"
)

func priorityChannel() {
	high := make(chan int)
	low := make(chan int)

	worker := func(tag string, task int) {
		fmt.Println(tag, task)
	}

	go func() {
		defer close(high)
		defer close(low)
		for {
			select {
			case v, ok := <-high:
				if !ok {
					return
				}
				worker("High", v)
			default:
				v, ok := <-low
				if !ok {
					return
				}
				worker("Low", v)
			}
		}
	}()

	for i := range 25 {
		if i%2 == 0 {
			high <- i
		} else {
			low <- i
		}
	}

	fmt.Println("Done")

}
