package module7

import "fmt"

func simpleRoutine() {
	done := make(chan bool)

	go func() {
		fmt.Println("A")
		done <- true
	}()
	go func() {
		fmt.Println("B")
		done <- true
	}()
	<-done
	<-done
}
