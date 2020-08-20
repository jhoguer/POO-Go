package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

func main() {
	c := course{name: "Go"}
	c.Print()
	fmt.Printf("El tipo es: %T\n", c)
}
