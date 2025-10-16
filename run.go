package main

import (
	"fmt"
	"learn/module1"
	"learn/module2"
	"learn/module7"
)

func main() {
	module1.PrintHello()
	// module2.FizzBuzz(50)
	x, y := module2.SwapVar(1, 2)
	fmt.Println("Var swap:", x, y)

	c := module2.C2F(24)
	fmt.Println("Temp converter:", c)

	sum := module2.SumOfArray([]int{1, 2, 3, 4, 5})
	fmt.Println("Sum of slice:", sum)

	maxInSlice := module2.MaxInSlice([]float64{1.5, 2.7, 0.3})
	fmt.Println("Max in slice:", maxInSlice)

	countOfTrue := module2.CountOfTrue([]bool{true, false, true})
	fmt.Println("count of bool:", countOfTrue)

	rev := module2.StringReverse("Hello")
	fmt.Println("String reverse:", rev)

	freqCounter := module2.MapFreqCounter("rohith")
	fmt.Println(freqCounter)

	oddEven := module2.OddEven(3)
	fmt.Println("3:", oddEven)

	grade := module2.GradeEvaluator(88)
	fmt.Println("Grade:", grade)

	fmt.Println("Sum of n:", module2.SumOfN(8))

	fmt.Println("Factorial:", module2.Factorial(5))

	sum, product := module2.SumAndProduct(2, 5)
	fmt.Println("Sum and product: ", sum, product)

	fmt.Println("Fib:", module2.Fibonacci(10))
	// 14
	{
		array := []string{"Apple", "Orange", "Banana", "Mango"}
		key := "Banana"
		index := module2.FindIndex(array, "Banana")
		fmt.Printf("%s in %v is %d\n", key, array, index)
	}
	// 15
	{
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := map[string]int{"b": 3, "c": 4}
		m3 := module2.MergeMap(m1, m2)
		fmt.Println(m3)
	}
	// 17
	fmt.Println("Sum of even:", module2.SumOfEven(10))
	// 18
	{
		m1 := map[string]int{"a": 1, "b": 2}
		out := module2.ReverseMap(m1)
		fmt.Println(m1, out)
	}
	// 19
	{
		prime := module2.IsPrime(5)
		nonPrime := module2.IsPrime(9)
		fmt.Println(prime, nonPrime)
	}
	// 20
	{
		slice := []int{1, 3, 5, 7, 9, 3, 23}
		out := module2.SliceAppend(slice, 11, 32, 54)
		fmt.Println(out)
	}

	// module3.Run()
	// module4.Run()
	// module5.Run()
	// module6.Run()
	module7.Run()
}
