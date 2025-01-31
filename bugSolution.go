```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Check if map is nil before accessing
    if m == nil {
        m = make(map[string]int)
    }
    m["a"] = 1
    fmt.Println(m["a"])

    //Alternative solution using blank assignment check
    var n map[string]int
    if val, ok := n["b"]; ok {
        fmt.Println(val)
    } else {
        n = make(map[string]int)
        n["b"] = 2
        fmt.Println(n["b"])
    }
}
```