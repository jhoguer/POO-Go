package main

import (
	"fmt"

	"github.com/jhoguer/POO-Go/tree/master/encapsulacion/course"
)

func main() {
	Go := course.New("Go desde cero", 11.78, false)

	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introdccion",
		2: "Estructuras",
		3: "POO",
	})

	Go.PrintClasses()

	Go.SetName("POO en Go")
	fmt.Println(Go.Name())
	fmt.Println(Go.UserIDs())
	// este metodo al ser NO exportado, no se puede usar fuera del paquete
	// Go.changePrice(34)

}
