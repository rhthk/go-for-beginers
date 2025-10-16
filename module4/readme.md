## 🧩 Go Structs and Methods — Practice Questions

This section covers Go structs, methods, and struct-based operations.
Each question includes a description and example to help you practice effectively.

### 1. Employee Salary Calculator

Description:
Define a struct Employee with fields Name, BaseSalary, and Bonus.
Write a method TotalSalary() that returns the total salary (base + bonus).
Demonstrate the use of this method by creating multiple employees and printing their total salaries.

Example:

```
emp := Employee{Name: "Alice", BaseSalary: 50000, Bonus: 5000}
fmt.Println(emp.TotalSalary()) // Output: 55000
```

### 2. Rectangle Area and Perimeter

Description:
Create a struct Rectangle with fields Width and Height.
Implement methods Area() and Perimeter() that return the respective values.
Write a program to take input for width and height and print both area and perimeter.

Example:

```
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area()) // Output: 50
fmt.Println(rect.Perimeter()) // Output: 30
```

### 3. Student Grade Updater

Description:
Define a struct Student with fields Name, Grade, and Marks.
Write a method UpdateMarks(newMarks int) that updates the student’s marks and adjusts the grade accordingly.
Use pointer receivers so the original struct is modified.

Example:

```
student := Student{Name: "Bob", Grade: "B", Marks: 70}
student.UpdateMarks(85)
fmt.Println(student.Grade) // Output: A
```

### 4. Complex Number Operations

Description:
Define a struct Complex with Real and Imag (float64).
Implement methods for:
• Add(c Complex) Complex
• Multiply(c Complex) Complex

Use these to perform addition and multiplication of two complex numbers.

Example:

```
a := Complex{Real: 3, Imag: 2}
b := Complex{Real: 1, Imag: 7}
sum := a.Add(b)
fmt.Println(sum) // Output: {4 9}
```

### 5. Library System — Struct Composition

Description:
Create two structs:
• Book → Title, Author, Price
• Library → contains a slice of Book

Write methods on Library to:
• AddBook(b Book)
• TotalValue() → returns the sum of all book prices

Demonstrate adding multiple books and printing total value.

Example:

```
lib := Library{}
lib.AddBook(Book{"Go Basics", "John", 300})
lib.AddBook(Book{"Go Advanced", "Jane", 500})
fmt.Println(lib.TotalValue()) // Output: 800
```
