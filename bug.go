```go
func main() {
    var m map[string]int
    m["a"] = 1 // This will panic if the map is nil
    fmt.Println(m["a"])
}
```