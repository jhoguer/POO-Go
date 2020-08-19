package main

import "fmt"

func main() {
	Go := &Course{
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

	fmt.Println("_________________________________________")
	Go.PrintClasses()
	fmt.Println("_________________________________________")
	Go.PrintName()
	fmt.Println("_________________________________________")

	// Go se encarga de gestionar si pasa el operador de desreferenciacion * o el operador de direccion &
	// (*Go).ChangePrice(15.42)

	// en este caso basta con solo usar el nombre de la instacia Go
	// sin necesidad de usar el operador * o & Go lo hace internamente
	Go.ChangePrice(17.42)
	fmt.Println(Go.Price)
	fmt.Println(Go)
	fmt.Println(*Go)
	fmt.Println(&Go)

}
