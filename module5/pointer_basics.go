package module5

import "fmt"

func pointerBasics() {
	i := 23
	p := &i
	fmt.Println(i, p)
}
