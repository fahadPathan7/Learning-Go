# <div align='center'>Float Formatting</div>

## Index
- [Float Formatting](#float-formatting)
  - [Index](#index)
  - [Float formatting](#float-formatting-1)

<br><br>

## Float formatting
Default float formatting is to print the float number with 6 decimal places. But we can change the number of decimal places by using the format specifiers.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%f\n", 3.141592653589793) // 3.141593
    fmt.Printf("%.2f\n", 3.141592653589793) // 3.14
    fmt.Printf("%.3f\n", 3.141592653589793) // 3.142
    fmt.Printf("%.4f\n", 3.141592653589793) // 3.1416
    fmt.Printf("%.5f\n", 3.141592653589793) // 3.14159
}
```

In the above code, we are using the format specifiers to print the float number with different decimal places.

<hr>

