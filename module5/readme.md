## Go Pointers & Interfaces – Practice Questions

### 1. Pointer Basics

Question:
Write a Go program to declare an integer variable and print its value and memory address using a pointer.
Goal: Understand how & and \* operators work in Go.

### 2. Pointer Dereferencing

Question:
Create a pointer to an integer, assign a value using the pointer, and print the updated integer.
Goal: Learn pointer dereferencing (\*ptr).

### 3. Nil Pointer Check

Question:
Declare a pointer without initialization and check if it’s nil.
Goal: Understand how uninitialized pointers behave in Go.

### 4. Function Parameter by Value vs Pointer

Question:
Write two functions — one that takes an integer by value, and another that takes a pointer to an integer.
Modify the integer in both and compare outputs.
Goal: Demonstrate difference between value and reference passing.

### 5. Swap Function using Pointers

Question:
Implement a function swap(a, b \*int) that swaps the values of two integers using pointers.
Goal: Practice modifying values through pointers in functions.

### 6. Return Pointer from Function

Question:
Write a function that returns a pointer to an integer initialized within the function.
Goal: Understand pointer return values and lifetime.

### 7. Pointer to Struct

Question:
Define a Person struct with name and age fields.
Create a pointer to a struct and modify its fields via pointer.
Goal: Practice using struct with pointers.

### 8. Function Receiving Struct Pointer

Question:
Write a function updateAge(p \*Person, newAge int) that updates the person’s age using pointer.
Goal: Learn how to pass struct pointers to functions.

### 9. Array and Pointer

Question:
Create an array of integers and use a pointer to access and modify its elements.
Goal: Explore pointer and array relationships.

### 10. Slice and Pointer Relationship

Question:
Explain and demonstrate why slices are reference types and how pointer usage differs with slices.
Goal: Understand slice and pointer interaction.

### 11. Interface and Pointer Receiver

Question:
Define an interface Shape with method Area() float64.
Implement it using a \*Circle struct (pointer receiver).
Goal: Learn how interfaces can be implemented using pointer receivers.

### 12. Interface with Value Receiver

Question:
Use the same Shape interface but implement it with a value receiver struct.
Compare behavior when assigning to interface.
Goal: Understand difference between pointer and value receiver in interface satisfaction.

### 13. Pointer to Interface vs Interface Holding Pointer

Question:
Write code showing the difference between a \*Shape and Shape that holds a pointer to a concrete type.
Goal: Clarify the conceptual difference between “pointer to interface” and “interface holding pointer”.
