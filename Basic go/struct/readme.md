# <div align='center'>Struct</div>

## Index
- [Struct](#struct)
  - [Index](#index)
  - [Struct in Go](#struct-in-go)
  - [Accessing Struct Fields](#accessing-struct-fields)
  - [Using Struct with Functions](#using-struct-with-functions)
  - [Methods with Pointer Receiver](#methods-with-pointer-receiver)

<br><br>

## Struct in Go
Struct is a user-defined type that contains a collection of named fields/properties. It is used to group related data together to form a single unit. It is similar to the class in object-oriented programming languages. Struct is a composite data type that can be used to group data of different types.

Struct is defined using the `type` keyword followed by the `struct` keyword. The fields of the struct are defined inside the curly braces `{}`. The fields are defined using the syntax `field_name data_type`.

```go
type struct_name struct {
    field_name1 data_type
    field_name2 data_type
    field_name3 data_type
    ...
}
```

<br><br>

## Accessing Struct Fields
The fields of the struct can be accessed using the dot `.` operator. The fields can be accessed using the syntax `struct_variable.field_name`.

```go
type Employee struct {
    name string
    age int
    salary float64
}

func main() {
    emp := Employee{name: "Kabir", age: 30, salary: 50000.50}
    fmt.Println("Name:", emp.name) // Output: Name: Kabir
    fmt.Println("Age:", emp.age) // Output: Age: 30
    fmt.Println("Salary:", emp.salary) // Output: Salary: 50000.50
}
```

<br><br>

## Using Struct with Functions
Struct can be used with functions to group related functions together. The functions that are associated with the struct are called methods. The methods are defined using the `func` keyword followed by the receiver type and the method name.

```go
type Employee struct {
    name string
    age int
    salary float64
}

func (e Employee) display() {
    fmt.Println("Name:", e.name)
    fmt.Println("Age:", e.age)
    fmt.Println("Salary:", e.salary)
}

func main() {
    emp := Employee{name: "Kabir", age: 30, salary: 50000.50}
    emp.display() // calls the associated method
}
```

In the above example, the `display` method is associated with the `Employee` struct. The method is called using the `struct_variable.method_name` syntax.

<br><br>

## Methods with Pointer Receiver
Methods can be defined with a pointer receiver. The pointer receiver is used to modify the struct fields inside the method. The pointer receiver is defined using the `*` symbol before the struct type.

```go
type Employee struct {
    name string
    age int
    salary float64
}

// without pointer receiver (no change in salary)
func (e Employee) salaryIncrement() {
    e.salary = e.salary + 5000
}

// with pointer receiver (change in salary)
func (e *Employee) updateSalary(newSalary float64) {
    e.salary = newSalary
}

func main() {
    emp := Employee{name: "Kabir", age: 30, salary: 50000.50}
    emp.salaryIncrement()
    fmt.Println("Updated Salary:", emp.salary) // Output: Updated Salary: 50000.50 (no change)
    emp.updateSalary(60000.75)
    fmt.Println("Updated Salary:", emp.salary) // Output: Updated Salary: 60000.75
}
```

<hr>