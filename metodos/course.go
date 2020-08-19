package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

// Como buena practica, si un metodo es tipo puntero *
// se aconseja ponerlos todos como tipo puntero
func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *Course) PrintName() {
	text := "El curso es: " + c.Name

	fmt.Println(text)
}

// Se pasa un puntero para poder modificar
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
