package main

import "fmt"

func successMessage() {
    fmt.Println("success")
}

func addNumbers(z, a int) int {
    return z+ a
}

func divideAndRemainder(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    successMessage()

    sum := addNumbers(3, 4)
    fmt.Println("Sum:", sum)

    quotient, remainder := divideAndRemainder(10, 3)
    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)
    
}