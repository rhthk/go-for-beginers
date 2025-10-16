package module4

import "fmt"

func Run() {
	// 1
	johnTotalSalary := employeeCalculator()
	fmt.Println("John total salary:", johnTotalSalary)
	// 2
	rect := Rectangle{length: 8, breadth: 4}
	fmt.Println(rect, rect.area(), rect.perimeter())
	// 3
	student := Student{name: "John", mark: 66, grade: "C"}
	fmt.Println(student)
	student.updateMark(88)
	fmt.Println(student)
	// 4
	complexA := Complex{real: 9, img: 3}
	complexB := Complex{real: 4, img: 6}
	complexA.addComplex(complexB)
	fmt.Println(complexA, complexB)
	// 5
	library := Library{}
	library.addBook(Book{name: "Go Basics", author: "Go lang", cost: 245})
	library.addBook(Book{name: "Go Intermediate", author: "Go lang, community", cost: 450})
	fmt.Println("Total cost of books:", library.totalValue())
}
