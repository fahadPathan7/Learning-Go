# <div align='center'>Loop</div>

## Index

<br><br>

## How Loop Works in Go
- A loop is a sequence of instructions that is continually repeated until a certain condition is reached.
- The for loop is used to execute a block of code multiple times.
- The for loop has three components: initialization, condition, and post statement.
- The initialization is executed only once at the beginning of the loop.
- The condition is evaluated before each iteration of the loop.
- The post statement is executed after each iteration of the loop.

<br><br>

## Loop Creation
The for loop is used to execute a block of code multiple times.

### Syntax
```go
for initialization; condition; post {
    // code
}
```

### Example
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

<br><br>

## Infinite Loop
An infinite loop is a loop that never stops. It continues to execute until the program is terminated.

### Syntax
```go
for {
    // code
}
```

### Example
```go
for {
    fmt.Println("Hello, World!")
}
```

<br><br>

## Loop Control Statements
Loop control statements are used to control the flow of a loop.

### Break Statement
The break statement is used to exit the loop.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break // exit the loop
    }
    fmt.Println(i)
}

output:
0
1
2
```

### Continue Statement
The continue statement is used to skip the current iteration of the loop.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        continue // skip the current iteration
    }
    fmt.Println(i)
}

output:
0
1
2
4
```

### Goto Statement
The goto statement is used to jump to a label in the code.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        goto end // jump to the end label
    }
    fmt.Println(i)
}

end:
fmt.Println("End of loop")

output:
0
1
2
End of loop
```

<br><br>

## Nested Loop
A nested loop is a loop inside another loop.

### Syntax
```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        // code
    }
}
```

### Example
```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Println(i, j)
    }
}

output:
0 0
0 1
0 2
1 0
1 1
1 2
2 0
2 1
2 2
```

<br><br>

## Looping Through Arrays and Slices
The range, len keyword is used to loop through arrays, slices, and maps.

### Syntax
```go
// using range
for index, value := range array {
    // code
}

// using len
for i := 0; i < len(array); i++ {
    // code
}
```

### Example
**using range:**
```go
var arr = [3]int{1, 2, 3}

for index, value := range arr {
    fmt.Println(index, value)
}

output:
0 1
1 2
2 3
```

**using len:**
```go
var arr = [3]int{1, 2, 3}

for i := 0; i < len(arr); i++ {
    fmt.Println(i, arr[i])
}

output:
0 1
1 2
2 3
```

<br><br>

## Looping Through Maps
The range keyword is used to loop through maps.

### Syntax
```go
for key, value := range map {
    // code
}
```

### Example
```go
var m = map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

for key, value := range m {
    fmt.Println(key, value)
}

output:
a 1
b 2
c 3
```

<br><br>

## Looping Through Strings
The range, len keyword is used to loop through strings.

### Syntax
```go
// using range
for index, value := range str {
    // code
}

// using len
for i := 0; i < len(str); i++ {
    // code
}
```

### Example
**using range:**
```go
var str = "hello"

for index, value := range str {
    fmt.Println(index, string(value))
}

output:
0 h
1 e
2 l
3 l
4 o
```

**using len:**
```go
var str = "hello"

for i := 0; i < len(str); i++ {
    fmt.Println(i, string(str[i]))
}

output:
0 h
1 e
2 l
3 l
4 o
```

<hr>

