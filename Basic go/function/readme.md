# <div align='center'>Function</div>

## Index
- [Function](#function)
  - [Index](#index)
  - [Function creation and calling](#function-creation-and-calling)
  - [Function with parameters](#function-with-parameters)
  - [Function with return type](#function-with-return-type)
  - [Function with named return values](#function-with-named-return-values)
  - [Function with variadic parameters](#function-with-variadic-parameters)
  - [Function with variadic parameters and normal parameters](#function-with-variadic-parameters-and-normal-parameters)
  - [Function can be of struct type](#function-can-be-of-struct-type)
  - [Function can be of interface type](#function-can-be-of-interface-type)
  - [Anonymous function](#anonymous-function)
    - [Anonymous function with parameters](#anonymous-function-with-parameters)
    - [Anonymous function with return type](#anonymous-function-with-return-type)
    - [Anonymous function assigning to a variable](#anonymous-function-assigning-to-a-variable)

<br><br>

## Function creation and calling

```go
// Function creation
func printHello() {
    fmt.Println("Hello")
}

func main() {
    printHello() // Function calling
}
```

<br><br>

## Function with parameters
parameters are the values that we pass to the function while calling it.

```go
func printHello(name string) {
    fmt.Println("Hello", name) // Hello Kamal
}

func main() {
    printHello("Kamal") // Function calling
}
```

parameters can be of different types

```go
func printHello(name string, age int) {
    fmt.Println("Hello", name, "your age is", age) // Hello Kamal your age is 25
}

func main() {
    printHello("Kamal", 25) // Function calling
}
```
here, we are passing two parameters, one is of type string and another is of type int.

<br><br>

## Function with return type
return type is the type of value that the function returns.

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(10, 20) // Function calling
    fmt.Println(result) // 30
}
```

function can return multiple values

```go
func addSub(a int, b int) (int, int) {
    return a + b, a - b
}

func main() {
    add, sub := addSub(10, 20) // Function calling
    fmt.Println(add, sub) // 30 -10
}
```

<br><br>

## Function with named return values
we can also return named values from the function.

```go
func addSub(a int, b int) (add int, sub int) {
    add = a + b
    sub = a - b
    return
}

func main() {
    add, sub := addSub(10, 20) // Function calling
    fmt.Println(add, sub) // 30 -10
}
```

<br><br>

## Function with variadic parameters
variadic parameters are the parameters that can take any number of values.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    result := sum(10, 20, 30, 40) // Function calling
    fmt.Println(result) // 100
}
```

<br><br>

## Function with variadic parameters and normal parameters
variadic parameters should be the **last parameter** in the function.

```go
func sum(a int, nums ...int) int {
    total := a
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    result := sum(10, 20, 30, 40) // Function calling
    fmt.Println(result) // 100
}
```

<br><br>

## Function can be of struct type
function can be of struct type.

```go
type person struct {
    name string
    age int
}

func (p person) printDetails() {
    fmt.Println("Name:", p.name, "Age:", p.age)
}

func main() {
    p := person{name: "Kamal", age: 25}
    p.printDetails() // Function calling
}
```

here, `printDetails` is a function of struct type `person`.

<br><br>

## Function can be of interface type
function can be of interface type.

```go
type person interface {
    printDetails()
}

type student struct {
    name string
    age int
}

func (s student) printDetails() {
    fmt.Println("Name:", s.name, "Age:", s.age)
}

func main() {
    var p person // Interface variable
    p = student{name: "Kamal", age: 25}
    p.printDetails() // Function calling
}
```

here, `printDetails` is a function of interface type `person`. We are creating a variable of interface type `person` and assigning a struct type `student` to it. Then we are calling the function `printDetails` using the interface variable.

<br><br>

## Anonymous function
anonymous functions are the functions without a name.

```go
func main() {
    func() {
        fmt.Println("Hello")
    }() // Function calling
}
```

### Anonymous function with parameters

```go
func main() {
    func(name string) {
        fmt.Println("Hello", name)
    }("Kamal") // Function calling
}
```

### Anonymous function with return type

```go
func main() {
    result := func(a int, b int) int {
        return a + b
    }(10, 20) // Function calling
    fmt.Println(result) // 30
}
```

### Anonymous function assigning to a variable

```go
func main() {
    add := func(a int, b int) int {
        return a + b
    }
    result := add(10, 20) // Function calling
    fmt.Println(result) // 30
}
```

---