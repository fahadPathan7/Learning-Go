# <div align='center'>Switch</div>

## Index
- [Switch](#switch)
  - [Index](#index)
  - [How Switch Works in Go](#how-switch-works-in-go)
  - [Switch Creation](#switch-creation)
    - [Syntax](#syntax)
    - [Example](#example)
  - [Switch Without Expression](#switch-without-expression)
    - [Syntax](#syntax-1)
    - [Example](#example-1)
  - [Fallthrough Statement](#fallthrough-statement)

<br><br>

## How Switch Works in Go
- The switch statement evaluates an expression and compares it with the values of each case.
- If the expression matches the value of a case, the corresponding block of code is executed.
- If no case matches the expression, the default block of code is executed.
- The default block of code is optional.
- The break statement is used to exit the switch statement.
- The fallthrough statement is used to execute the next case block of code.
- The switch statement can also be used without an expression.
- The switch statement without an expression is used to execute the first case that evaluates to true. (This is similar to the if-else-if statement.)

<br><br>

## Switch Creation
The switch statement is used to execute one block of code from multiple blocks of code based on a condition.

### Syntax
```go
switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}
```

### Example
```go
var day = "Monday"

switch day {
case "Monday":
    fmt.Println("Today is Monday")
case "Tuesday":
    fmt.Println("Today is Tuesday")
case "Wednesday":
    fmt.Println("Today is Wednesday")
case "Thursday":
    fmt.Println("Today is Thursday")
case "Friday":
    fmt.Println("Today is Friday")
case "Saturday":
    fmt.Println("Today is Saturday")
case "Sunday":
    fmt.Println("Today is Sunday")
default:
    fmt.Println("Invalid day")
}
```

<br><br>

## Switch Without Expression
The switch statement without an expression is used to execute the first case that evaluates to true.

### Syntax
```go
switch {
case condition1:
    // code
case condition2:
    // code
default:
    // code
}
```

### Example
```go
var num = 10

switch {
case num > 0:
    fmt.Println("Number is positive")
case num < 0:
    fmt.Println("Number is negative")
default:
    fmt.Println("Number is zero")
}
```

<br><br>

## Fallthrough Statement
The fallthrough statement is used to execute the next case block of code.

```go
var num = 10

switch {
case num > 0:
    fmt.Println("Number is positive")
    fallthrough
case num < 0:
    fmt.Println("Number is negative")
default:
    fmt.Println("Number is zero")
}
```

In the above example, the output will be:
```
Number is positive
Number is negative
```

The fallthrough statement is not recommended because it makes the code less readable and harder to maintain. It is better to use the break statement to exit the switch statement.

<hr>