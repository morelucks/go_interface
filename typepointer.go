package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) increment() {
    c.value++ 
}

func main() {
    counter := Counter{value: 10}
    counter.increment() 
    fmt.Println("Counter value:", counter.value) 
}
