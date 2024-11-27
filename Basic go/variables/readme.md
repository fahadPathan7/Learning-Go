# <div align='center'> Variables </div>

## Index
- [ Variables ](#-variables-)
  - [Index](#index)
  - [Variables types](#variables-types)
  - [Variable Data Types](#variable-data-types)
  - [Declaring variables](#declaring-variables)
  - [Constants](#constants)
  - [Type Conversion](#type-conversion)
  - [Zero Values](#zero-values)
  - [Scope of Variables](#scope-of-variables)

<br><br>

## Variables types
In Go, there are three types of variables:
- **Local variables**: These are the variables that are declared inside a function. They are only accessible within the function.
- **Global variables**: These are the variables that are declared outside a function. They are accessible throughout the program.
- **Block variables**: These are the variables that are declared inside a block of code. They are only accessible within that block.

<br><br>

## Variable Data Types
In Go, variables can be of different data types. The following are the basic data types in Go:
- **Integers**: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- **Floating-point numbers**: float32, float64
- **Complex numbers**: complex64, complex128
- **Boolean**: bool
- **Strings**: string

<br><br>

## Declaring variables
In Go, variables are declared using the `var` keyword. The syntax for declaring a variable is as follows:
```go
var variable_name data_type
```

Here is an example of declaring a variable:
```go
var age int
```

In the above example, we have declared a variable named `age` of type `int`.

We can also declare multiple variables in a single line:
```go
var name, age string
```

In the above example, we have declared two variables named `name` and `age` of type `string`.

Variables can also be declared and initialized at the same time:
```go
var name string = "Fahad"
```

In the above example, we have declared a variable named `name` of type `string` and initialized it with the value `"Fahad"`.

We can also omit the data type and let Go infer it:
```go
var name = "Fahad"
```

In the above example, we have declared a variable named `name` and initialized it with the value `"Fahad"`. Go will infer the data type of the variable based on the value.

We can also use the short variable declaration `:=` to declare and initialize a variable:
```go
name := "Fahad"
```

In the above example, we have declared a variable named `name` and initialized it with the value `"Fahad"` using the short variable declaration.

<br><br>

## Constants
In Go, constants are declared using the `const` keyword. The syntax for declaring a constant is as follows:
```go
const constant_name data_type = value
```

Here is an example of declaring a constant:
```go
const pi float64 = 3.14159
```

In the above example, we have declared a constant named `pi` of type `float64` with the value `3.14159`.

Constants can also be declared without specifying the data type:
```go
const pi = 3.14159
```

In the above example, we have declared a constant named `pi` and initialized it with the value `3.14159`. Go will infer the data type of the constant based on the value.

Constants can also be declared in a group:
```go
const (
    pi = 3.14159
    e = 2.71828
)
```

In the above example, we have declared two constants named `pi` and `e` and initialized them with the values `3.14159` and `2.71828` respectively.

<br><br>

## Type Conversion
In Go, we can convert variables from one type to another using type conversion. The syntax for type conversion is as follows:
```go
new_variable_name := data_type(variable_name)
```

Here is an example of type conversion:
```go
var num int = 10
var result float64 = float64(num)

fmt.Println(result) // 10.0
```

In the above example, we have converted the variable `num` of type `int` to a variable `result` of type `float64`.

<br><br>

## Zero Values
In Go, variables that are declared without an explicit initialization are given a zero value. The zero value of a variable depends on its data type. The following are the zero values for different data types:

- **Integers**: 0
- **Floating-point numbers**: 0.0
- **Complex numbers**: 0+0i
- **Boolean**: false
- **Strings**: ""
- **Pointers**: nil
- **Functions**: nil
- **Interfaces**: nil
- **Slices**: nil

Here is an example of zero values:
```go
var num int
var result float64
var flag bool
var name string

fmt.Println(num)    // 0
fmt.Println(result) // 0.0
fmt.Println(flag)   // false
fmt.Println(name)   // ""
```

In the above example, we have declared variables `num`, `result`, `flag`, and `name` without initializing them. The zero values of these variables are printed using the `fmt.Println()` function.

<br><br>

## Scope of Variables
In Go, the scope of a variable is the region of the program where the variable is accessible. The scope of a variable depends on where it is declared. The following are the different scopes of variables in Go:
- **Local scope**: Variables declared inside a function have local scope and are only accessible within that function.
    ```go
    package main

    import "fmt"

    func main() {
        var name string = "Fahad"

        fmt.Println(name) // Fahad
    }
    ```
    In the above example, we have declared a variable named `name` inside the `main()` function. The variable `name` is accessible only within the `main()` function.
- **Global scope**: Variables declared outside a function have global scope and are accessible throughout the program.
    ```go
    package main

    import "fmt"

    var name string = "Fahad"

    func main() {
        fmt.Println(name) // Fahad
    }
    ```
    In the above example, we have declared a variable named `name` outside the `main()` function. The variable `name` is accessible within the `main()` function.
- **Block scope**: Variables declared inside a block of code have block scope and are only accessible within that block.
    ```go
    package main

    import "fmt"

    func main() {
        if true {
            var name string = "Fahad"
            fmt.Println(name) // Fahad
        }
    }
    ```
    In the above example, we have declared a variable named `name` inside the `if` block. The variable `name` is accessible only within the `if` block.

<hr>