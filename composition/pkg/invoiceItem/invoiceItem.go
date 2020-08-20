package invoiceItem

// Item contains of information of a invoiceItem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// nuevo tipo personalizado
type Items []Item

// funcion constructora del nuevo tipo Items que sacamos de Item
func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}

	return is
}

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}

	return total
}
