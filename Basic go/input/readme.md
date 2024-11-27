# <div align='center'> Input </div>

## Index
- [ Input ](#-input-)
  - [Index](#index)
  - [Input](#input)

<br><br>

## Input
In Go, the `fmt` package provides functions to read input from the user. The `Scanf` function is used to read input from the user. The syntax for reading input using the `Scanf` function is as follows:
```go
fmt.Scanf("format", &variable_name)
```

Here is an example of reading input from the user:
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Println("Enter your name:")
    fmt.Scanf("%s", &name)
    fmt.Println("Hello,", name)
}
```

In the above example, we have declared a variable named `name` of type `string`. We prompt the user to enter their name using the `fmt.Println()` function. We then read the input using the `fmt.Scanf()` function with the format specifier `%s` and store the input in the `name` variable. Finally, we print the message "Hello, " followed by the user's name.

The `Scanf` function can be used with different format specifiers to read different types of input. The following are some commonly used format specifiers:

- `%d`: Reads an integer value.
- `%f`: Reads a floating-point value.
- `%s`: Reads a string value.
- `%c`: Reads a single character.
- `%t`: Reads a boolean value.
- `%v`: Reads any value.
- `%p`: Reads a pointer value.
- `%T`: Reads the type of the value.

The `Scanf` function reads input until a newline character is encountered. If we want to read input until a space character is encountered, we can use the `%q` format specifier.

```go
var name string
fmt.Scanf("%q", &name)
```

In the above example, the `Scanf` function will read input until a space character is encountered and store it in the `name` variable.

The `Scanf` function can also be used to read multiple values from the user. We can specify multiple format specifiers in the format string and pass the addresses of the variables where we want to store the input.

```go
var name string
var age int
fmt.Scanf("%s %d", &name, &age)
```

In the above example, the `Scanf` function will read a string value followed by an integer value from the user and store them in the `name` and `age` variables, respectively.

The `Scanf` function returns the number of items successfully scanned. We can use this return value to check if the input was read successfully.

```go
var name string
var age int
n, err := fmt.Scanf("%s %d", &name, &age)
if err != nil {
    fmt.Println("Error reading input:", err)
} else {
    fmt.Println("Input read successfully:", n, "items scanned")
}
```

In the above example, we use the `n` variable to store the number of items scanned by the `Scanf` function. If an error occurs while reading input, the `err` variable will be non-nil, and we can print an error message.

Inputs can be scanned without any format specifier using the `Scan` function. The `Scan` function reads input as a string and returns the number of bytes read.

```go
var input, name string
fmt.Scan(&input, &name)
```

In the above example, the `Scan` function will read two strings from the user and store them in the `input` and `name` variables.

<hr>
