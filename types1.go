package main

import "fmt"

//struct
type Person struct {
    name string
    age  int
}

// value receiver
func (p Person) greet() {
    fmt.Println("Hello, my name is", p.name)
}

func main() {
    person := Person{name: "Alice", age: 25}
    person.greet()
}
