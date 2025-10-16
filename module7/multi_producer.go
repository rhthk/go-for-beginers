package module7

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	value int
	id    int
}

func multipleProducer() {
	var wg sync.WaitGroup
	const timeout = 500 * time.Millisecond
	producers := [3]chan int{}
	workerDone := make(chan struct{})
	for i := range producers {
		producers[i] = make(chan int)
	}

	in := make(chan Item)

	wg.Add(len(producers))
	for i, prod := range producers {
		go func(id int, ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				in <- Item{value: v, id: id}
			}
		}(i, prod)
	}

	selector := func() {
		defer close(workerDone)
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		for {
			select {
			case it, ok := <-in:
				if !ok {
					return
				}
				fmt.Println("Data received", it.id, it.value)
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(timeout)
			case <-timer.C:
				fmt.Println("maintenance")
				timer.Reset(timeout)
			}
		}
	}
	go selector()

	// --- Simulate producers sending values at different times ---

	for n := range 3 {
		go func(id int) {
			defer close(producers[id%3])

			for i := range 5 {
				producers[id%3] <- i
				// Make sleep time depend on goroutine id and iteration
				sleepTime := time.Duration((id+1)*(i*10)*100) * time.Millisecond
				time.Sleep(sleepTime)
			}
		}(n)
	}

	// Wait for all forwarders to finish, then close fan-in channel
	waiter := func() {
		wg.Wait()
		close(in)
	}
	go waiter()

	// Wait until worker exits
	<-workerDone
	fmt.Println("all done")
}
