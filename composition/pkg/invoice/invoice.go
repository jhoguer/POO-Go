package invoice

import (
	"github.com/jhoguer/POO-Go/tree/composition/composition/pkg/customer"
	"github.com/jhoguer/POO-Go/tree/composition/composition/pkg/invoiceItem"
)

// Invpice is the struct of a invoice/factura
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer  // Relacion 1 a 1
	items   []invoiceItem.Item // Relacion de 1 a muchos
}

// New return a new invoice
func New(country, city string, client customer.Customer, items []invoiceItem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
