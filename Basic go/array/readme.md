# <div align='center'> Go Array Details </div>

## Index
- [ Go Array Details ](#-go-array-details-)
  - [Index](#index)
  - [Array Declaration](#array-declaration)
    - [1D Array](#1d-array)
    - [2D Array](#2d-array)
  - [Array Operations](#array-operations)
    - [Accessing and Updating Elements](#accessing-and-updating-elements)
    - [Length of the Array and Iterating Over the Array](#length-of-the-array-and-iterating-over-the-array)
    - [Deleting Elements](#deleting-elements)

<br><br>

## Array Declaration

### 1D Array

```go
var arr [5]int // [0, 0, 0, 0, 0]
arr := [5]int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
arr := [...]int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
arr := [5]int {} // [0, 0, 0, 0, 0]
```

- In the first example, we declare an array of size 5 with all elements initialized to 0.
- In the second example, we declare an array of size 5 with elements 1, 2, 3, 4, and 5.
- In the third example, we declare an array of size 5 with elements 1, 2, 3, 4, and 5. The size of the array is inferred from the number of elements.
- In the fourth example, we declare an array of size 5 with no elements initialized.

### 2D Array

```go
var arr [2][3]int // [[0, 0, 0], [0, 0, 0]]
arr := [2][3]int{{1, 2, 3}, {4, 5, 6}} // [[1, 2, 3], [4, 5, 6]]
arr := [...][3]int{{1, 2, 3}, {4, 5, 6}} // [[1, 2, 3], [4, 5, 6]]
arr := [2][3]int {} // [[0, 0, 0], [0, 0, 0]]
```

- In the first example, we declare a 2D array of size 2x3 with all elements initialized to 0. (by default)
- In the second example, we declare a 2D array of size 2x3 with elements 1, 2, 3, 4, 5, and 6.
- In the third example, we declare a 2D array of size 2x3 with elements 1, 2, 3, 4, 5, and 6. The size of the array is inferred from the number of elements.
- In the fourth example, we declare a 2D array of size 2x3 with no elements initialized.

<br><br>

## Array Operations

### Accessing and Updating Elements

```go
arr := [5]int{1, 2, 3, 4, 5}

// Accessing elements
fmt.Println(arr[0]) // 1
fmt.Println(arr[1]) // 2

// Updating elements
arr[0] = 10
fmt.Println(arr[0]) // 10
```

In the above example, we declare an array of size 5 with elements 1, 2, 3, 4, and 5. We access the elements using the index and update the elements using the index.

### Length of the Array and Iterating Over the Array

```go
arr := [5]int{1, 2, 3, 4, 5}

// Length of the array
fmt.Println(len(arr)) // 5

// Iterating over the array
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

We find the length of the array using the `len()` function and iterate over the array using a for loop.

### Deleting Elements

In Go, we cannot delete elements from an array. However, we can create a new array without the element we want to delete.

```go
arr := [5]int{1, 2, 3, 4, 5}
index := 2

newArr := [4]int{}
for i := 0; i < len(arr); i++ {
    if i < index {
        newArr[i] = arr[i]
    } else if i > index {
        newArr[i-1] = arr[i]
    }
}

fmt.Println(newArr) // [1, 2, 4, 5]
```

In the above example, we declare an array of size 5 with elements 1, 2, 3, 4, and 5. We want to delete the element at index 2. We create a new array without the element at index 2.

<hr>


