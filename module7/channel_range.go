package module7

import "fmt"

func rangeOverChannel() {
	ch := make(chan int, 10)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("from channel:", v)
	}
}
