package main

import "fmt"

func main() {

    fruits := map[string]string{
        "apple":  "red",
        "banana": "yellow",
        "grape":  "purple",
    }

    fmt.Println("Color of an apple:", fruits["apple"])
    fmt.Println("Color of a banana:", fruits["banana"])
    fmt.Println("Color of a grape:", fruits["grape"])

    fruits["orange"] = "orange"



    color, exists := fruits["jackfruit"]
    if exists {
        fmt.Println("Color of a jackfruit:", color)
    } else {
        fmt.Println("Color of a jackfruit is not available.")
    }


    delete(fruits, "grape")

    for fruit, color := range fruits {
        fmt.Println("Fruit:", fruit, "| Color:", color)
    }
}