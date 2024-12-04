# <div align='center'>Pointer</div>

## Index
- [Pointer](#pointer)
  - [Index](#index)
  - [Pointer](#pointer-1)
  - [Pointer to Pointer](#pointer-to-pointer)
  - [Struct can be passed as a pointer](#struct-can-be-passed-as-a-pointer)
  - [Pointer to Array](#pointer-to-array)

<br><br>

## Pointer
Pointer is a variable that stores the memory address of another variable.

```go
func main() {
    var a int = 10
    var p *int // p is a pointer to an integer

    p = &a // p stores the memory address of a

    fmt.Println("Value of a:", a) // 10
    fmt.Println("Address of a:", &a) // 0xc0000b6010
    fmt.Println("Value of p:", p) // 0xc0000b6010
    fmt.Println("Value of *p:", *p) // 10
}
```

- `&` is used to get the memory address of a variable.
- `*` is used to get the value of a pointer.
- `*` is also used to declare a pointer variable.

<br><br>

## Pointer to Pointer
Pointer to pointer is a pointer variable that stores the memory address of another pointer variable.

```go
func main() {
    var a int = 10
    var p *int // p is a pointer to an integer
    var q **int // q is a pointer to a pointer to an integer

    p = &a // p stores the memory address of a
    q = &p // q stores the memory address of p

    fmt.Println("Value of a:", a) // 10
    fmt.Println("Address of a:", &a) // 0xc0000b6010
    fmt.Println("Value of p:", p) // 0xc0000b6010
    fmt.Println("Value of *p:", *p) // 10
    fmt.Println("Value of q:", q) // 0xc0000b6020
    fmt.Println("Value of *q:", *q) // 0xc0000b6010
    fmt.Println("Value of **q:", **q) // 10
}
```

- `**` is used to get the value of a pointer to pointer.
- `**` is also used to declare a pointer to pointer variable.

<br><br>

## Struct can be passed as a pointer
Struct can be passed as a pointer to a function.

```go
type Person struct {
    Name string
    Age int
}

func main() {
    p := Person{"John", 25}
    fmt.Println("Before:", p) // {John 25}
    changeName(&p)
    fmt.Println("After:", p) // {Alice 25}
}

func changeName(p *Person) {
    p.Name = "Alice"
}
```

Here, `changeName` function takes a pointer to a `Person` struct. So, the changes made inside the function will reflect outside the function.

<br><br>

## Pointer to Array
Pointer to an array can be passed to a function.

```go
func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println("Before:", arr) // [1 2 3]
    changeArray(&arr)
    fmt.Println("After:", arr) // [4 5 6]
}

func changeArray(arr *[3]int) {
    for i := 0; i < len(arr); i++ {
        arr[i] += 3
    }
}
```

Here, `changeArray` function takes a pointer to an array of integers. So, the changes made inside the function will reflect outside the function.

<hr>