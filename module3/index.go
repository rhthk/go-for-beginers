package module3

import "fmt"

func Run() {
	// 1
	fmt.Println(sum()(5, 3))
	// 2
	fmt.Println(calculate(add, 3, 6))
	fmt.Println(calculate(multiply, 3, 6))
	// 3
	doubler := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println(doubler(5), triple(5))
	// 4
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	// 5
	adder := sumClosure()
	fmt.Println(adder(5))
	fmt.Println(adder(10))
	fmt.Println(adder(3))
}
