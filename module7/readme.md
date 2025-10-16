# Easy (5)

### 1. Buffered vs unbuffered channel send/receive

Show difference between buffered and unbuffered channels: which operations block and when a goroutine waits.

example: input and expected output

```go
func main() {
ch1 := make(chan int) // unbuffered
ch2 := make(chan int, 1) // buffered size 1

    go func() { ch1 <- 1; fmt.Println("sent ch1") }()
    time.Sleep(10 * time.Millisecond)
    go func() { ch2 <- 2; fmt.Println("sent ch2") }()

    fmt.Println(<-ch1)
    fmt.Println(<-ch2)

}
```

expected output

```
sent ch1 1
sent ch2 2
```

Explanation: send to unbuffered ch1 synchronizes with receiver; sending to buffered ch2 may not block because capacity allows one value.

### 2. Range over a channel and closing it

Use close so a receiver using for v := range ch exits the loop.

example: input and expected output

// input

```go
func main() {
ch := make(chan int)
go func() {
for i := 1; i <= 3; i++ { ch <- i }
close(ch)
}()
for v := range ch {
fmt.Println(v)
}
}
```

expected output

```
1
2
3
```

### 3. Simple goroutine ordering (race nondeterminism)

Demonstrate that goroutine scheduling is nondeterministic; prints may appear in either order.

example: input and expected output

// input

```go
func main() {
done := make(chan bool)
go func() { fmt.Println("A"); done <- true }()
go func() { fmt.Println("B"); done <- true }()
<-done
<-done
}
```

expected output (one possible)

```
A
B
```

possible alternative

```
B
A
```

### 4. Select with default (non-blocking send/receive)

Use select with a default case to perform non-blocking sends/receives.

example: input and expected output

```go
// input
func main() {
ch := make(chan int, 1)
ch <- 1
select {
case ch <- 2:
fmt.Println("sent 2")
default:
fmt.Println("channel full, skipping send")
}
select {
case v := <-ch:
fmt.Println("received", v)
default:
fmt.Println("nothing to receive")
}
}
```

expected output

```
channel full, skipping send
received 1
```

### 5. Closing a channel vs sending nil value

Closing signals “no more values”; receivers get zero value after channel drained and ok==false.

example: input and expected output

// input

```go
func main() {
ch := make(chan int)
go func() { ch <- 10; close(ch) }()
v, ok := <-ch
fmt.Println(v, ok)
v2, ok2 := <-ch
fmt.Println(v2, ok2)
}
```

expected output

```
10 true
0 false
```

---

## Hard (6)

### 6. Fan-in: merging multiple channels into one with goroutines and select

Implement a merge that reads from multiple input channels and forwards values to a single output channel until all inputs are closed.

example: input and expected output

```
// input (conceptual)
a := make(chan int)
b := make(chan int)
out := merge(a,b) // merge implemented using goroutines+select
// a sends 1,2 then close; b sends 10,20 then close
```

expected output (order nondeterministic between sources)

```
1
10
2
20
```

(Any interleaving of values from a and b is valid; all four values must appear once.)

### 7. Worker pool with context cancellation and graceful shutdown

Design workers reading tasks from a channel; stop accepting new tasks using context.Context cancellation and wait for in-flight tasks to finish.

example: input and expected output

```
// input (conceptual)
ctx, cancel := context.WithCancel(context.Background())
taskCh := make(chan int)
go startWorkers(ctx, 3, taskCh) // startWorkers spawns 3 workers
taskCh <- 1; taskCh <- 2
cancel() // request shutdown
// no more tasks sent
```

expected output

```
worker processed 1
worker processed 2
// then program exits with all workers returned (no leaks)
```

(Exact worker IDs may vary; purpose is graceful stop without losing in-flight tasks.)

### 8. Detecting and avoiding deadlocks with buffered channels and goroutines

Given code that deadlocks when a goroutine blocks on send because no receiver exists, modify to avoid deadlock (e.g., add buffer, start receiver first, or use select with timeout).

example: input and expected output

```
// problematic input
func main() {
ch := make(chan int) // unbuffered
ch <- 1 // deadlock: no receiver and main goroutine blocks
}
```

expected output (after fix, e.g., add goroutine receiver)

```
// if fixed with a receiver goroutine:
go func(){ fmt.Println(<-ch) }()
ch <- 1
// prints:
1
```

### 9. Select across many channels with timeout and fallbacks

Implement a select loop that consumes from many producers, but if nothing arrives for N milliseconds perform a periodic maintenance task.

example: input and expected output

```
// input (conceptual)
producers := []chan int{p1,p2,p3}
loop with select { case <-time.After(100\*time.Millisecond): doMaintenance() }
```

expected output

```
When producers idle for >100ms, "maintenance" is printed periodically.
When producers send, values printed as received.
```

### 10. Using select to implement prioritized channel handling

How to implement priority (prefer channel A over B) using select without busy-waiting.

example: input and expected output

```go
// input (conceptual)
select {
    case v := <-highPriority: process(v)
    default:
        select {
            case v := <-lowPriority: process(v)
            case <-time.After(...): ...
        }
}
```

expected output

```
If highPriority has data, it's always processed first. If not, lowPriority may be processed.
```

### 11. Closing a channel while multiple goroutines write — coordination pattern

Explain safe patterns for shutting down writers (e.g., only producer closes channel, or use separate “done” channel to signal writers to stop) and show a code pattern.

example: input and expected output

```
// input (conceptual)
producers check a quit channel and stop sending; only captain goroutine closes the results channel once all done
```

expected output

```
No panic: values received until closed, then range exits cleanly.
```

---

Hard — Edge case scenarios (3)

### 12. Sending on a closed channel (panic) — reproduce and explain

Show a program where a goroutine sends into a channel that gets closed by another goroutine, causing a runtime panic; explain how to prevent it.

example: input and expected output

```go
// input
func main() {
ch := make(chan int)
go func() {
ch <- 1 // send
}()
go func() {
close(ch)
}()
time.Sleep(10 \* time.Millisecond)
}
```

expected output

```
panic: send on closed channel
// stack trace...
```

(Prevention: ensure single owner closes channel or use a stop signal rather than closing from multiple places.)

### 13. select with a nil channel: branch permanently disabled (edge case)

When a channel variable is nil, any case <-nil or case nil <- in select is never selectable — useful for dynamically enabling/disabling cases.

example: input and expected output

```go
// input
func main() {
a := make(chan int)
var b chan int // nil
go func(){ a <- 5 }()
select {
case v := <-a:
fmt.Println("a:", v)
case v := <-b:
fmt.Println("b:", v)
case <-time.After(50\*time.Millisecond):
fmt.Println("timeout")
}
}
```

expected output

```
a: 5
```

(case <-b is disabled because b is nil. If a had no send, select would choose timeout.)

### 14. Goroutine leak via time.After in long-running loop (edge case)

Repeatedly using time.After inside a loop without proper cancellation can create many unused timers (memory leak). Show safer alternative using a single time.Ticker or time.NewTimer reuse.

example: input and expected output

// input (problematic)

```go
for {
select {
case v := <-ch:
fmt.Println(v)
case <-time.After(time.Hour):
// intended periodic check (but creates timer each loop)
}
}
// safer: use ticker := time.NewTicker(time.Hour); defer ticker.Stop()
// then use case <-ticker.C:
```

expected output (behavioral)

```
Safer version uses one timer/ticker; program won't leak timers if loop iterates frequently.
```
