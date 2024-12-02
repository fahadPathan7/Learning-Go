# <div align='center'>Condition</div>

## Index
- [Condition](#condition)
  - [Index](#index)
  - [Condition](#condition-1)
    - [If Statement](#if-statement)
    - [If-Else Statement](#if-else-statement)
    - [If-Else-If Statement](#if-else-if-statement)
    - [Nested If Statement](#nested-if-statement)

<br><br>

## Condition
Condition is a statement that evaluates to true or false. It is used to make decisions based on the value of a variable or an expression.

### If Statement
The if statement is used to execute a block of code only if a condition is true.

```go
var num = 10

if num > 0 {
    fmt.Println("Number is positive")
}
```

### If-Else Statement
The if-else statement is used to execute a block of code if the condition is true and another block of code if the condition is false.

```go
var num = 10

if num > 0 {
    fmt.Println("Number is positive")
} else {
    fmt.Println("Number is negative")
}
```

### If-Else-If Statement
The if-else-if statement is used to execute one block of code from multiple blocks of code based on a condition.

```go
var num = 10

if num > 0 {
    fmt.Println("Number is positive")
} else if num < 0 {
    fmt.Println("Number is negative")
} else {
    fmt.Println("Number is zero")
}
```

### Nested If Statement
The nested if statement is an if statement inside another if statement.

```go
var num = 10

if num >= 0 {
    if num == 0 {
        fmt.Println("Number is zero")
    } else {
        fmt.Println("Number is positive")
    }
} else {
    fmt.Println("Number is negative")
}
```

<hr>