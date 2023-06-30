package main

import "fmt"

func main() {
    
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    for key, value := range colors {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }
    
}