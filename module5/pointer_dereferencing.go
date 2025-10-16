package module5

import "fmt"

func pointerDereferencing() {
	i := 25
	p := &i
	fmt.Println("Before:", i, p)
	*p = 45
	fmt.Println("After:", i, p)
}
