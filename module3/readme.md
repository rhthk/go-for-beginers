# 🧠 Go Programming Exercises: Function Values, Closures, and CLI Arguments

This set of exercises focuses on Go functions as first-class values, closures, and reading command-line arguments.

---

### 1. Function as a Value

**Description:**  
Write a Go program where you assign a function that adds two numbers to a variable and call it through that variable.

**Example:**  
Input:

```
3 5
```

Output:

```
Sum: 8
```

---

### 2. Passing Function as Argument

**Description:**  
Create a function `operate(a, b int, f func(int, int) int)` that applies the given function `f` to `a` and `b`. Use it to compute sum, product, and difference.

**Example Output:**

```
Sum: 8
Product: 15
Difference: 2
```

---

### 3. Returning Functions (Closures)

**Description:**  
Write a function `makeMultiplier(factor int)` that returns another function which multiplies any given number by that factor.

**Example:**

```go
double := makeMultiplier(2)
fmt.Println(double(5)) // Output: 10
```

---

### 4. Counter Closure

**Description:**  
Implement a closure that returns a function to count how many times it has been called.

**Example:**

```go
next := counter()
fmt.Println(next()) // 1
fmt.Println(next()) // 2
fmt.Println(next()) // 3
```

---

### 5. Sum Closure

**Description:**  
Write a closure that keeps a running total of numbers passed to it.

**Example:**

```go
sum := makeAdder()
sum(5)  // 5
sum(10) // 15
sum(3)  // 18
```

---

### 6. CLI Argument Reader

**Description:**  
Write a program that prints all command-line arguments passed to it.

**Example Command:**

```
go run main.go apple banana orange
```

**Output:**

```
Args: [apple banana orange]
```

---

### 7. CLI Calculator

**Description:**  
Create a simple CLI calculator. The program should take three CLI arguments: two integers and an operator (+, -, \*, /).

**Example Command:**

```
go run main.go 10 + 5
```

**Output:**

```
Result: 15
```

---

### 8. CLI Greeter

**Description:**  
Write a program that accepts a name from CLI arguments and prints a greeting message.

**Example Command:**

```
go run main.go Alice
```

**Output:**

```
Hello, Alice!
```

---

### 9. Closure for Filtering

**Description:**  
Write a closure that returns a function to filter numbers greater than a given threshold.

**Example:**

```go
filter := greaterThan(10)
fmt.Println(filter([]int{5, 12, 8, 20})) // [12 20]
```

---

### 10. Closure with CLI Integration

**Description:**  
Create a closure `repeatPrinter` that takes a string and returns a function that prints it a specified number of times.  
Accept the string and number from CLI arguments.

**Example Command:**

```
go run main.go hello 3
```

**Output:**

```
hello
hello
hello
```

---

✅ **Topics Covered**

- Function variables
- Passing and returning functions
- Closures for maintaining state
- Reading command-line arguments via `os.Args`
- Combining closures with CLI input
