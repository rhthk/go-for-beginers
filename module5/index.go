package module5

import "fmt"

func Run() {
	fmt.Println("Module 5")
	// 1
	pointerBasics()
	// 2
	pointerDereferencing()
	// 3
	pointerNil()
	// 4
	n := 4
	fmt.Println(n, passByValue(n), n)
	fmt.Println(n, passByPointer(&n), n)
	// 5
	var a, b int = 4, 5
	swapByPointer(&a, &b)
	fmt.Println(a, b)
	// 6
	fmt.Println(*returnPointer())
	// 7
	person := &Person{name: "John Doe", age: 23}
	fmt.Println(*person)
	(*person).age = 25
	fmt.Println(*person)
	// 8
	updateAge(person, 27)
	fmt.Println(*person)
	// 9
	var slice = []int{1, 2, 3, 4, 5}
	p := &slice
	(*p)[0] = 88
	// 10
	fmt.Println(slice)
	modifySlice(slice)
	fmt.Println(slice)
	// 11
	var circle Shape = &Circle{r: 3}
	fmt.Println("Area:", circle.area())
	// 12
	fmt.Println("Perimeter:", circle.perimeter())
}
