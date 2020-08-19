package main

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

	Go.PrintClasses()

}
