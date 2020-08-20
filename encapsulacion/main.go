package main

import (
	"github.com/jhoguer/POO-Go/tree/master/encapsulacion/course"
)

func main() {
	Go := &course.Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Clases: map[uint]string{
			1: "Introdccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()

	// este metodo al ser NO exportado, no se puede usar fuera del paquete
	// Go.changePrice(34)

}
