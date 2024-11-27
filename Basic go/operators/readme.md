# <div align='center'> Operators </div>

## Index
- [ Operators ](#-operators-)
  - [Index](#index)
  - [Operators](#operators)
  - [Arithmetic Operators](#arithmetic-operators)
  - [Relational Operators](#relational-operators)
  - [Logical Operators](#logical-operators)
  - [Bitwise Operators](#bitwise-operators)
  - [Assignment Operators](#assignment-operators)
  - [Miscellaneous Operators](#miscellaneous-operators)

<br><br>

## Operators
Operators are special symbols that are used to perform operations on variables and values. In Go, there are several types of operators, such as arithmetic operators, relational operators, logical operators, etc. The following are the different types of operators in Go:

- **Arithmetic Operators**
- **Relational Operators**
- **Logical Operators**
- **Bitwise Operators**
- **Assignment Operators**
- **Miscellaneous Operators**

<br><br>

## Arithmetic Operators
Arithmetic operators are used to perform mathematical operations on variables and values. The following are the arithmetic operators in Go:

- **Addition (+)**: Adds two operands.
    ```go
        a := 10
        b := 20
        c := a + b
        fmt.Println(c) // Output: 30
    ```
- **Subtraction (-)**: Subtracts the second operand from the first operand.
    ```go
        a := 20
        b := 10
        c := a - b
        fmt.Println(c) // Output: 10
    ```
- **Multiplication (*)**: Multiplies two operands.
    ```go
        a := 10
        b := 20
        c := a * b
        fmt.Println(c) // Output: 200
    ```
- **Division (/)**: Divides the first operand by the second operand.
    ```go
        a := 20
        b := 10
        c := a / b
        fmt.Println(c) // Output: 2
    ```
- **Modulus (%)**: Returns the remainder of the division of the first operand by the second operand.
    ```go
        a := 22
        b := 10
        c := a % b
        fmt.Println(c) // Output: 2
    ```
- **Increment (++)**: Increases the value of the operand by 1. (does not supoort prefix increment)
    ```go
        a := 10
        a++
        fmt.Println(a) // Output: 11
    ```
- **Decrement (--)**: Decreases the value of the operand by 1. (does not supoort prefix decrement)
    ```go
        a := 10
        a--
        fmt.Println(a) // Output: 9
    ```

<br><br>

## Relational Operators
Relational operators are used to compare two values. The following are the relational operators in Go:

- **Equal to (==)**: Checks if the values of two operands are equal.
    ```go
        a := 10
        b := 20
        fmt.Println(a == b) // Output: false
    ```
- **Not equal to (!=)**: Checks if the values of two operands are not equal.
    ```go
        a := 10
        b := 20
        fmt.Println(a != b) // Output: true
    ```
- **Greater than (>)**: Checks if the value of the first operand is greater than the value of the second operand.
    ```go
        a := 20
        b := 10
        fmt.Println(a > b) // Output: true
    ```
- **Less than (<)**: Checks if the value of the first operand is less than the value of the second operand.
    ```go
        a := 10
        b := 20
        fmt.Println(a < b) // Output: true
    ```
- **Greater than or equal to (>=)**: Checks if the value of the first operand is greater than or equal to the value of the second operand.
    ```go
        a := 20
        b := 10
        fmt.Println(a >= b) // Output: true
    ```
- **Less than or equal to (<=)**: Checks if the value of the first operand is less than or equal to the value of the second operand.
    ```go
        a := 10
        b := 10
        fmt.Println(a <= b) // Output: true
    ```

<br><br>

## Logical Operators
Logical operators are used to combine multiple conditions. The following are the logical operators in Go:

- **Logical AND (&&)**: Returns true if both operands are true.
    ```go
        a := true
        b := false
        fmt.Println(a && b) // Output: false
    ```
- **Logical OR (||)**: Returns true if at least one of the operands is true.
    ```go
        a := true
        b := false
        fmt.Println(a || b) // Output: true
    ```
- **Logical NOT (!)**: Returns true if the operand is false and vice versa.
    ```go
        a := true
        fmt.Println(!a) // Output: false
    ```
<br><br>

## Bitwise Operators
Bitwise operators are used to perform bitwise operations on integer values. The following are the bitwise operators in Go:

- **Bitwise AND (&)**: Performs a bitwise AND operation on the binary representations of the two operands. (if all bits are 1, then 1, else 0)
    ```go
        a := 5 // 101
        b := 3 // 011
        c := a & b // 001
        fmt.Println(c) // Output: 1
    ```
- **Bitwise OR (|)**: Performs a bitwise OR operation on the binary representations of the two operands. (if any bit is 1, then 1, else 0)
    ```go
        a := 5 // 101
        b := 3 // 011
        c := a | b // 111
        fmt.Println(c) // Output: 7
    ```
- **Bitwise XOR (^)**: Performs a bitwise XOR (exclusive OR) operation on the binary representations of the two operands. (if bits are different, then 1, else 0)
    ```go
        a := 5 // 101
        b := 3 // 011
        c := a ^ b // 110
        fmt.Println(c) // Output: 6
    ```
- **Bitwise NOT (~)**: Performs a bitwise NOT operation on the binary representation of the operand. (inverts the bits)
    ```go
        a := 5 // 101
        c := ^a // 010
        fmt.Println(c) // Output: -6
    ```
- **Left Shift (<<)**: Shifts the bits of the first operand to the left by the number of positions specified by the second operand.
    ```go
        a := 5 // 101
        c := a << 1 // 1010
        fmt.Println(c) // Output: 10
    ```
- **Right Shift (>>)**: Shifts the bits of the first operand to the right by the number of positions specified by the second operand.
    ```go
        a := 5 // 101
        c := a >> 1 // 10
        fmt.Println(c) // Output: 2
    ```

<br><br>

## Assignment Operators
Assignment operators are used to assign values to variables. The following are the assignment operators in Go:

- **Assignment (=)**: Assigns the value of the right operand to the left operand.
    ```go
        a := 10
        b := 20
        a = b
        fmt.Println(a) // Output: 20
    ```
- **Addition Assignment (+=)**: Adds the value of the right operand to the left operand and assigns the result to the left operand.
    ```go
        a := 10
        b := 20
        a += b // Equivalent to a = a + b
        fmt.Println(a) // Output: 30
    ```
- **Subtraction Assignment (-=)**: Subtracts the value of the right operand from the left operand and assigns the result to the left operand.
    ```go
        a := 20
        b := 10
        a -= b // Equivalent to a = a - b
        fmt.Println(a) // Output: 10
    ```
- **Multiplication Assignment (*=)**: Multiplies the value of the right operand with the left operand and assigns the result to the left operand.
    ```go
        a := 10
        b := 20
        a *= b // Equivalent to a = a * b
        fmt.Println(a) // Output: 200
    ```
- **Division Assignment (/=)**: Divides the value of the left operand by the right operand and assigns the result to the left operand.
    ```go
        a := 20
        b := 10
        a /= b // Equivalent to a = a / b
        fmt.Println(a) // Output: 2
    ```
- **Modulus Assignment (%=)**: Calculates the remainder of the division of the left operand by the right operand and assigns the result to the left operand.
    ```go
        a := 22
        b := 10
        a %= b // Equivalent to a = a % b
        fmt.Println(a) // Output: 2
    ```
- **Left Shift Assignment (<<=)**: Shifts the bits of the left operand to the left by the number of positions specified by the right operand and assigns the result to the left operand.
    ```go
        a := 5 // 101
        a <<= 1 // Equivalent to a = a << 1
        fmt.Println(a) // Output: 10
    ```
- **Right Shift Assignment (>>=)**: Shifts the bits of the left operand to the right by the number of positions specified by the right operand and assigns the result to the left operand.
    ```go
        a := 5 // 101
        a >>= 1 // Equivalent to a = a >> 1
        fmt.Println(a) // Output: 2
    ```
- **Bitwise AND Assignment (&=)**: Performs a bitwise AND operation on the binary representations of the left and right operands and assigns the result to the left operand.
    ```go
        a := 5 // 101
        b := 3 // 011
        a &= b // Equivalent to a = a & b
        fmt.Println(a) // Output: 1
    ```
- **Bitwise OR Assignment (|=)**: Performs a bitwise OR operation on the binary representations of the left and right operands and assigns the result to the left operand.
    ```go
        a := 5 // 101
        b := 3 // 011
        a |= b // Equivalent to a = a | b
        fmt.Println(a) // Output: 7
    ```
- **Bitwise XOR Assignment (^=)**: Performs a bitwise XOR operation on the binary representations of the left and right operands and assigns the result to the left operand.
    ```go
        a := 5 // 101
        b := 3 // 011
        a ^= b // Equivalent to a = a ^ b
        fmt.Println(a) // Output: 6
    ```
- **Bitwise NOT Assignment (&=)**: Performs a bitwise NOT operation on the binary representation of the left operand and assigns the result to the left operand.
    ```go
        a := 5 // 101
        a ^= a // Equivalent to a = ^a
        fmt.Println(a) // Output: -6
    ```

<br><br>

## Miscellaneous Operators
There are a few other operators in Go that are used for specific purposes:

- **Address Operator (&)**: Returns the memory address of a variable.
    ```go
        a := 10
        fmt.Println(&a) // Output: 0xc0000b6010
    ```
- **Pointer Operator (*)**: Declares a pointer variable or dereferences a pointer variable.
    ```go
        a := 10
        b := &a
        fmt.Println(*b) // Output: 10
    ```

<hr>