package module7

import "fmt"

func selectBasic() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("send 2")
	default:
		fmt.Println("Channel full")
	}
	select {
	case v := <-ch:
		fmt.Println("received", v)
	default:
		fmt.Println("nothing to receive")
	}
}
