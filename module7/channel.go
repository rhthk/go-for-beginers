package module7

import (
	"fmt"
	"time"
)

func print(ch chan int, n int) {
	fmt.Println("Sent...", n)
	ch <- n
}

func channelBasic() {
	ch1 := make(chan int)
	ch2 := make(chan int, 2)
	go print(ch1, 1)
	time.Sleep(100 * time.Microsecond)
	go print(ch2, 2)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
