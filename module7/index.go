package module7

import "fmt"

func Run() {
	// 1
	// channelBasic()
	// // 2
	// rangeOverChannel()
	// // 3
	// simpleRoutine()
	// // 4
	// selectBasic()
	// // 5
	// closeChannel()
	// // 6
	// channelMerge()
	// // 7
	// workerPoolMaster()
	// // 8
	// deadlock()
	// 9
	// multipleProducer()
	// 10
	// priorityChannel()
}

func closeChannel() {
	ch := make(chan int)
	go func() { ch <- 10; close(ch) }()
	v, ok := <-ch
	fmt.Println(v, ok)
	v2, ok := <-ch
	fmt.Println(v2, ok)

}

func channelMerge() {
	// make buffered channels (so we can send before the receiver starts)
	a := make(chan int, 3)
	b := make(chan int, 3)

	// populate a and close it
	for i := range 3 {
		a <- i
	}
	close(a)

	// populate b and close it
	for i := range 3 {
		b <- i + 3
	}
	close(b)
	// merge reads from two input channels and returns a channel that will be closed
	merge := func(x, y <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			// keep reading while either x or y is non-nil
			for x != nil || y != nil {
				select {
				case v, ok := <-x:
					if !ok {
						// x closed — stop selecting from it
						x = nil
						continue
					}
					out <- v
				case v, ok := <-y:
					if !ok {
						// y closed — stop selecting from it
						y = nil
						continue
					}
					out <- v
				}
			}
		}()
		return out
	}
	for v := range merge(a, b) {
		fmt.Println("Out:", v)
	}
}
