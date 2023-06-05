package main

import "fmt"

func main() {
   
    age := 25
    switch {
    case age < 18:
        fmt.Println("You are under 18.")
    case age >= 18 && age < 65:
        fmt.Println("You are an adult.")
    case age >= 65:
        fmt.Println("You are a senior citizen.")
    }
}