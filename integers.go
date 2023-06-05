package main

import "fmt"

func main() {
    var a int = 10
    var b int8 = 4
    var c int16 = 200
    var d int32 = -1000
    var e int64 = 9876543210

    sum := a + int(b)
    difference := c - int16(d)
    product := int64(b) * e

    fmt.Println("a + b =", sum)
    fmt.Println("c - d =", difference)
    fmt.Println("b * e =", product)
}
