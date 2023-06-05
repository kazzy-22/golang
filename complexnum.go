package main

import "fmt"

func main() {
    a := complex(2.222, 3)
    b := complex(1.05, -2.5)
    sum := a + b
    difference := a - b
    product := a * b
	
    fmt.Println("a + b =", sum)
    fmt.Println("a - b =", difference)
    fmt.Println("a * b =", product)

    realPart := real(sum)
    imagPart := imag(sum)

    fmt.Println("Real part of a + b =", realPart)
    fmt.Println("Imaginary part of a + b =", imagPart)
}