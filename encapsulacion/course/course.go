package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
