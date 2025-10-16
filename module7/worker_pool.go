package module7

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range jobs {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Worker", id, task)
	}
}

func startWorkers(ctx context.Context, n int, in <-chan int) <-chan struct{} {
	jobs := make(chan int)
	var wg sync.WaitGroup
	tracker := func() {
		defer close(jobs)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context closed")
				return
			case task, ok := <-in:
				fmt.Println("Received task:", task)
				if !ok {
					return
				}
				jobs <- task
			}
		}
	}

	go tracker()

	wg.Add(n)
	for i := range n {
		go worker(i+1, jobs, &wg)
	}

	done := make(chan struct{})

	waiter := func() {
		wg.Wait()
		close(done)
	}

	go waiter()

	return done
}

func workerPoolMaster() {
	fmt.Println("Start master")
	ctx, cancel := context.WithCancel(context.Background())
	taskCh := make(chan int)

	done := startWorkers(ctx, 2, taskCh)

	go func() {
		for i := range 10 {
			time.Sleep(50 * time.Millisecond)
			taskCh <- i + 1
		}
		close(taskCh)
	}()

	time.Sleep(3000 * time.Millisecond)
	cancel()
	<-done
	fmt.Println("Worker canceled")
}

/*
Main
  |
  | startWorkers(ctx, 3, taskCh)
  v
startWorkers
  |
  |-- create jobs chan
  |-- create wg
  |
  |-- spawn forwarder goroutine -------------------------
  |   forwarder: loop {
  |     select {
  |     case <-ctx.Done():       // stop forwarding, defer close(jobs)
  |       return
  |     case t, ok := <-in:      // read from external taskCh
  |       if !ok { return }      // input closed -> return
  |       select {
  |       case jobs <- t:        // forward task into internal jobs
  |       case <-ctx.Done():     // cancel while trying to forward
  |         return
  |       }
  |     }
  |   }
  |   (defer close(jobs))
  |
  |-- spawn N worker goroutines --------------------------
  |   for i := 0 .. n-1 {
  |     wg.Add(1)
  |     go worker(id=i+1) {
  |       defer wg.Done()
  |       for t := range jobs {   // loop until jobs closed & drained
  |         process(t)
  |       }
  |       // worker returns
  |     }
  |   }
  |
  |-- spawn waiter goroutine ----------------------------
  |   go func() {
  |     wg.Wait()        // block until all wg.Done() called
  |     close(done)      // signal to caller that workers are finished
  |   }()
  |
  |-- return done <-chan struct{}
  v
Main (continues)
  |
  |-- task producer goroutine: send tasks into taskCh; then close(taskCh)
  |
  |-- after short sleep: cancel()  // signals ctx.Done()
  |
  |-- <-done    // wait until done channel closed (all workers exited)
  |
  `-- program exits
*/
