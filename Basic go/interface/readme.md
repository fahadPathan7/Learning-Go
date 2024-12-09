# <div align='center'>Interface in Go</div>

## Index
- [Interface in Go](#interface-in-go)
  - [Index](#index)
  - [Introduction](#introduction)
  - [Implementing an Interface](#implementing-an-interface)
  - [Interface with Multiple Methods](#interface-with-multiple-methods)
  - [Interface with Embedded Interface](#interface-with-embedded-interface)
  - [Interface as value in map](#interface-as-value-in-map)

<br><br>

## Introduction
An interface is a collection of method signatures that a type can implement. It is a way to achieve polymorphism in Go. An interface is defined using the `interface` keyword. An interface can be implemented by any type that has the methods defined in the interface.

<br><br>

## Implementing an Interface
To implement an interface, a type must define all the methods defined in the interface. The type can have additional methods, but it must have all the methods defined in the interface.

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape // Declare a variable of type Shape
    s = Circle{Radius: 5} // Assign a Circle type to the Shape variable
    fmt.Println(s.Area()) // Output: 78.5
}
```

In the above example, we have defined an interface `Shape` with a method `Area()`. We have defined a struct `Circle` with a field `Radius`. We have implemented the `Area()` method for the `Circle` struct. We have assigned a `Circle` type to the `Shape` variable `s`. We have called the `Area()` method using the `Shape` variable `s`.

<br><br>

## Interface with Multiple Methods
An interface can have multiple methods. A type must define all the methods defined in the interface to implement the interface.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func main() {
    var s Shape
    s = Circle{Radius: 5}
    fmt.Println(s.Area()) // Output: 78.5
    fmt.Println(s.Perimeter()) // Output: 31.4
}
```

In the above example, we have defined an interface `Shape` with two methods `Area()` and `Perimeter()`. We have defined a struct `Circle` with a field `Radius`. We have implemented the `Area()` and `Perimeter()` methods for the `Circle` struct. We have assigned a `Circle` type to the `Shape` variable `s`. We have called the `Area()` and `Perimeter()` methods using the `Shape` variable `s`.

<br><br>

## Interface with Embedded Interface
An interface can embed another interface. A type must define all the methods defined in the embedded interface to implement the interface.

```go
type Shape interface {
    Area() float64
}

type Perimeter interface {
    Perimeter() float64
}

type Geometry interface {
    Shape
    Perimeter
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func main() {
    var g Geometry
    g = Circle{Radius: 5}
    fmt.Println(g.Area()) // Output: 78.5
    fmt.Println(g.Perimeter()) // Output: 31.4
}
```

In the above example, we have defined an interface `Shape` with a method `Area()`. We have defined an interface `Perimeter` with a method `Perimeter()`. We have defined an interface `Geometry` that embeds the `Shape` and `Perimeter` interfaces. We have defined a struct `Circle` with a field `Radius`. We have implemented the `Area()` and `Perimeter()` methods for the `Circle` struct. We have assigned a `Circle` type to the `Geometry` variable `g`. We have called the `Area()` and `Perimeter()` methods using the `Geometry` variable `g`.

<br><br>

## Interface as value in map
An interface can be used as a value in a map. It is used to store different types of values in a map.

```go
func main() {
    m := make(map[string]interface{})
    m["name"] = "John"
    m["age"] = 30
    m["salary"] = 50000.0
    fmt.Println(m) // Output: map[age:30 name:John salary:50000]
}
```

In the above example, we have created a map `m` with a key of type `string` and a value of type `interface{}`. We have stored different types of values in the map `m`.

<hr>

