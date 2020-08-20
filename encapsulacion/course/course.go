package course

import "fmt"

// Nuestra estructura es exportada
// los campos de la estructura son exportados
type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	clases  map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}

func (c *course) IsFree() bool {
	return c.isFree
}

func (c *course) SetUserIDs(userID []uint) {
	c.userIDs = userID
}

func (c *course) UserIDs() []uint {
	return c.userIDs
}

func (c *course) SetName(name string) {
	c.name = name
}

func (c *course) Name() string {
	return c.name
}

func (c *course) SetClasses(classes map[uint]string) {
	c.clases = classes
}

// Los metodos tambien son exportados
func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
