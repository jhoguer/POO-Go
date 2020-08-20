package main

import (
	"fmt"

	"github.com/jhoguer/POO-Go/tree/composition/composition/pkg/customer"
	"github.com/jhoguer/POO-Go/tree/composition/composition/pkg/invoice"
	"github.com/jhoguer/POO-Go/tree/composition/composition/pkg/invoiceItem"
)

func main() {
	cliente := customer.New("Jhon Guerrero", "Calle 17 # 26 - 17", "45352355")
	curso1 := invoiceItem.New(1, "Curso de Go", 17.55)
	curso2 := invoiceItem.New(2, "Curso de Bases de datos con Go", 20.54)
	curso3 := invoiceItem.New(3, "Curso de POO con Go", 12.34)
	curso4 := invoiceItem.New(4, "API con Go", 25.90)

	factura := invoice.New(
		"Colombia",
		"Puerto Asis",
		cliente,
		[]invoiceItem.Item{
			curso1,
			curso2,
			curso3,
			curso4,
		},
	)

	i := invoice.New(
		"Colombia",
		"Popayan",
		customer.New("Jhon Guerrero", "Calle 17 # 26 - 17", "45352355"),
		[]invoiceItem.Item{
			invoiceItem.New(1, "Curso de Go", 17.55),
			invoiceItem.New(2, "Curso de Bases de datos con Go", 20.54),
			invoiceItem.New(3, "Curso de POO con Go", 12.34),
			invoiceItem.New(4, "API con Go", 30.05),
		},
	)

	fmt.Printf("%+v", i)
	i.SetTotal()
	fmt.Printf("%+v", i)

	fmt.Println("_____________________________________")

	fmt.Printf("%+v", factura)
	i.SetTotal()
	fmt.Printf("%+v", factura)

}
