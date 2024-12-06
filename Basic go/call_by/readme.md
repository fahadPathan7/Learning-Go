# <div align='center'>Call By</div>

## Index
- [Call By](#call-by)
  - [Index](#index)
  - [Call By Value](#call-by-value)
  - [Call By Reference](#call-by-reference)
  - [Call By Value vs Call By Reference](#call-by-value-vs-call-by-reference)

<br><br>

## Call By Value
In call by value, the `value` of the actual parameter is passed to the formal parameter. The changes made to the formal parameter have no effect on the actual parameter.

```go
func changeValue(x int) {
    x = 10
}

func main() {
    x := 5
    changeValue(x) // value will not change actually
    fmt.Println(x) // Output: 5
}
```

<br><br>

## Call By Reference
In call by reference, the `address` of the actual parameter is passed to the formal parameter. The changes made to the formal parameter have an effect on the actual parameter.

```go
func changeValue(x *int) {
    *x = 10
}

func main() {
    x := 5
    changeValue(&x) // value will change actually
    fmt.Println(x) // Output: 10
}
```

<br><br>

## Call By Value vs Call By Reference

| Call By Value | Call By Reference |
| --- | --- |
| The value of the actual parameter is passed to the formal parameter. | The address of the actual parameter is passed to the formal parameter. |
| The changes made to the formal parameter have no effect on the actual parameter. | The changes made to the formal parameter have an effect on the actual parameter. |
| It is used when we don't want to change the actual parameter. | It is used when we want to change the actual parameter. |

<hr>