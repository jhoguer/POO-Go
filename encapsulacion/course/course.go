package course

import "fmt"

// Nuestra estructura es exportada
// los campos de la estructura son exportados
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

// Los metodos tambien son exportados
func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

// Este metodo es NO exportado
func (c *Course) changePrice(price float64) {
	c.Price = price
}
