package course

import "fmt"

// Nuestra estructura es exportada
// los campos de la estructura son exportados
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

// Los metodos tambien son exportados
func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

// Este metodo es NO exportado
func (c *course) changePrice(price float64) {
	c.Price = price
}
