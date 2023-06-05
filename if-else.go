package main

import "fmt"

func main() {
    age := 22
    if age <= 18 {
        fmt.Println("You are eligible to drive.")
    } else {
        fmt.Println("You are not eligible to drive.")
    }
}