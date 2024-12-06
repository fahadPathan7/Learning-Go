# <div align='center'>Map in Go</div>

## Index

<br><br>

## Map
A map is a collection of key-value pairs. It is an unordered collection. It is used to store data in the form of pairs. The key is unique in the map. The key is used to retrieve the value. The key and value can be of any type.

```go
package main

import "fmt"

func main() {
    // Create a map
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    // Access the value using key
    fmt.Println(m["one"]) // Output: 1

    // Add a new key-value pair
    m["four"] = 4
    fmt.Println(m) // Output: map[four:4 one:1 three:3 two:2]

    // Delete a key-value pair
    delete(m, "two")
    fmt.Println(m) // Output: map[four:4 one:1 three:3]
}
```

<br><br>

## Iterating over a map
We can iterate over a map using a `for` loop. The `range` keyword is used to iterate over the map.

```go
func main() {
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    for k, v := range m {
        fmt.Println(k, v) // key, value
    }
}
```

<br><br>

## Passing a map to a function
A map can be passed to a function. The changes made to the map in the function have an effect on the actual map.

```go
func modifyMap(m map[string]int) {
    m["one"] = 10
}

func main() {
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    modifyMap(m)
    fmt.Println(m) // map[one:10 three:3 two:2]
}
```

<br><br>

## adding a key-value pair to a map
A key-value pair can be added to a map using the key.

```go
func main() {
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    m["four"] = 4
    fmt.Println(m) // map[four:4 one:1 three:3 two:2]
}
```

<br><br>

## Deleting a key-value pair from a map
A key-value pair can be deleted from a map using the key.

```go
func main() {
    m := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    delete(m, "two")
    fmt.Println(m) // map[one:1 three:3]
}
```

<hr>