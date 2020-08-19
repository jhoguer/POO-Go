package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func main() {
	Go := Course{
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

	// Tambien podemos declarar estructuras sin usar los nombres
	// de cada campo, solo si se declaran todos los campos y en
	// el orden en que se definieron
	NodeJs := Course{
		"NodeJs desde cero",
		15.34,
		false,
		[]uint{23, 12, 56, 16, 89},
		map[uint]string{
			1: "Fundamentos de Js",
			2: "Asincronismo",
			3: "Express",
			4: "PassportJs",
		},
	}

	css := Course{Name: "CSS desde cero", IsFree: true}

	js := Course{}
	js.Name = "Curso de JS"
	js.UserIDs = []uint{12, 67}

	fmt.Println(Go.Name)
	fmt.Println(NodeJs.Name)
	fmt.Printf("%+v\n\n", css)
	fmt.Printf("%+v\n\n", js)

	fmt.Println("_________________________________________")

	Go.PrintClasses()
	fmt.Println("_________________________________________")
	js.PrintClasses()

	js.Clases = map[uint]string{
		1: "Basicos Js",
		2: "Avanzados Js",
	}
	js.PrintClasses()
	fmt.Println("_________________________________________")

	NodeJs.PrintClasses()
}
