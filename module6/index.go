package module6

import (
	"cmp"
	"fmt"
)

func Run() {
	fmt.Println("Module 6: Generic types")
	// 1
	fmt.Println(min(10, 20))
	fmt.Println(min(3.14, 2.71))
	fmt.Println(min("apple", "banana"))
	// 2
	nums := []int{1, 2, 3, 4, 5, 6}
	even := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(even)
	// 3
	intBox := Box[int]{value: 3}
	strBox := Box[string]{value: "Three"}
	fmt.Println(intBox.get(), strBox.get())

}

func min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func filter[T any](items []T, predicate func(T) bool) []T {
	newList := make([]T, 0)
	for _, v := range items {
		if isOk := predicate(v); isOk {
			newList = append(newList, v)
		}
	}
	return newList
}

type Box[T any] struct {
	value T
}

func (box *Box[T]) get() T {
	return box.value
}
