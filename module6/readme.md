### 1. Generic Function to Find Minimum Value

Topic: Generic Function with Type Constraints
Description:
Write a generic function Min[T constraints.Ordered](a, b T) T that takes two values of any ordered type (int, float64, string, etc.) and returns the smaller one.

Example:
Input:

```
fmt.Println(Min(10, 20))
fmt.Println(Min(3.14, 2.71))
fmt.Println(Min("apple", "banana"))
```

Expected Output:

```
10
2.71
apple
```

### 2. Generic Slice Filter Function

Topic: Generic Function with Function Parameters
Description:
Write a generic function Filter[T any](items []T, predicate func(T) bool) []T that returns a new slice containing only elements that satisfy the predicate.

Example:
Input:

```
nums := []int{1, 2, 3, 4, 5, 6}
even := Filter(nums, func(n int) bool { return n%2 == 0 })
fmt.Println(even)
```

Expected Output:

```
[2 4 6]
```

### 3. Generic Type with Methods

Topic: Generic Struct with Type Parameter
Description:
Create a generic struct Box[T any] that can hold a value of any type. Implement a method Get() that returns the stored value.

Example:

Input:

```
intBox := Box[int]{value: 42}
strBox := Box[string]{value: "Hello, Go!"}

fmt.Println(intBox.Get())
fmt.Println(strBox.Get())
```

Expected Output:

```
42
Hello, Go!
```
