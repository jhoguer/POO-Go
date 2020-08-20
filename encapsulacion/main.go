package main

import (
	"fmt"

	"github.com/jhoguer/POO-Go/tree/master/encapsulacion/course"
)

func main() {
	Go := course.New("Go desde cero", 11.78, false)

	Go.UserIDs = []uint{12, 56, 89}
	Go.Clases = map[uint]string{
		1: "Introdccion",
		2: "Estructuras",
		3: "Maps",
	}

	Go.PrintClasses()
	fmt.Printf("%+v", Go)

	// este metodo al ser NO exportado, no se puede usar fuera del paquete
	// Go.changePrice(34)

}
