# <div align='center'>Slice</div>

## Index
- [Slice](#slice)
  - [Index](#index)
  - [Slice](#slice-1)
  - [Creating a slice](#creating-a-slice)
  - [Length and capacity](#length-and-capacity)
  - [What if the capacity is not provided?](#what-if-the-capacity-is-not-provided)
  - [What if the length exceeds the capacity?](#what-if-the-length-exceeds-the-capacity)
  - [Modifying a slice](#modifying-a-slice)
  - [Slicing a slice](#slicing-a-slice)
  - [Appending elements to a slice](#appending-elements-to-a-slice)
  - [Removing elements from a slice](#removing-elements-from-a-slice)
  - [Copying a slice](#copying-a-slice)
  - [Iterating over a slice](#iterating-over-a-slice)
  - [Passing a slice to a function](#passing-a-slice-to-a-function)
  - [Passing a slice to a function by reference](#passing-a-slice-to-a-function-by-reference)

<br><br>

## Slice
A slice is a data structure in Go that provides a more powerful and flexible way to work with sequences of elements compared to arrays.
- A slice is a reference to a contiguous segment of an array.
- Size of a slice is not fixed.
- Removing or adding elements to a slice is easy.

<br><br>

## Creating a slice
A slice can be created using the `make` function or by using a slice literal.
```go
// Using make function
slice := make([]int, 5, 10) // Type, length, capacity

// Using slice literal
slice := []int{1, 2, 3, 4, 5}

// Creating a slice from an array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:3] // [2, 3]
```

<br><br>

## Length and capacity
- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

```go
slice := make([]int, 5, 10)
fmt.Println(len(slice)) // 5
fmt.Println(cap(slice)) // 10
```

<br><br>

## What if the capacity is not provided?
If the capacity is not provided, the capacity of the slice will be equal to the length of the slice.
```go
slice := make([]int, 5)
fmt.Println(len(slice)) // 5
fmt.Println(cap(slice)) // 5
```

<br><br>

## What if the length exceeds the capacity?
If the length exceeds the capacity, a new array will be created with double the capacity of the previous array.
```go
slice := make([]int, 5, 5)
slice = append(slice, 6)
fmt.Println(len(slice)) // 6
fmt.Println(cap(slice)) // 10
```

<br><br>

## Modifying a slice
- Modifying a slice modifies the underlying array.
- Modifying a slice does not affect other slices that share the same underlying array.

```go
slice := []int{1, 2, 3, 4, 5}
slice[0] = 10
fmt.Println(slice) // [10, 2, 3, 4, 5]
```

<br><br>

## Slicing a slice
A slice can be sliced to create a new slice.
```go
slice := []int{1, 2, 3, 4, 5}
newSlice := slice[1:3] // [2, 3]
```

<br><br>

## Appending elements to a slice
Elements can be appended to a slice using the `append` function.
```go
slice := []int{1, 2, 3, 4, 5}
slice = append(slice, 6)
fmt.Println(slice) // [1, 2, 3, 4, 5, 6]
```

<br><br>

## Removing elements from a slice
Elements can be removed from a slice using the `append` function.
```go
slice := []int{1, 2, 3, 4, 5}
slice = append(slice[:2], slice[3:]...) // [1, 2, 4, 5]
```

... means that the slice is being expanded into individual elements. It is called the variadic parameter. It is used to pass a variable number of arguments to a function. In this case, it is used to pass the elements of the slice.

<br><br>

## Copying a slice
A slice can be copied to create a new slice.
```go
slice := []int{1, 2, 3, 4, 5}
newSlice := make([]int, len(slice))
copy(newSlice, slice) // first argument is the destination, second argument is the source
```

<br><br>

## Iterating over a slice
A slice can be iterated over using a `for` loop.
```go
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    fmt.Println(i, v) // index, value
}

// another way
for i := 0; i < len(slice); i++ {
    fmt.Println(slice[i])
}
```

here, range returns the index and the value of the element at that index.

<br><br>

## Passing a slice to a function
A slice can be passed to a function.
```go
func modifySlice(slice []int) {
    // it can not change the length and capacity of the slice
    slice[0] = 10
}

slice := []int{1, 2, 3, 4, 5}
modifySlice(slice)
fmt.Println(slice) // [10, 2, 3, 4, 5]
```

<br><br>

## Passing a slice to a function by reference
A slice can be passed to a function by reference.
```go
func modifySlice(slice *[]int) {
    // it can change the length and capacity of the slice
    (*slice)[0] = 10
}

slice := []int{1, 2, 3, 4, 5}
modifySlice(&slice)
fmt.Println(slice) // [10, 2, 3, 4, 5]
```

<hr>


