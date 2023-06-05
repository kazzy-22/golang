package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Location string
}

func main() {
	person := Person{
		Name:     "jonny liver",
		Age:      42,
		Location: "nalasupara",
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Location:", person.Location)

	person.Name = "vijay malya"
	person.Age = 52
	person.Location = "dhoka nagar"

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Location:", person.Location)
}