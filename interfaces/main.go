package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

// Embebiendo 2 interfaces
type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adios soy %s\n", p.Name)
}

// Implementando una interface propia de Go
func (p Person) String() string {
	return "Hola soy una persona y mi nombre es: " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Adios soy %s\n", t)
}

// Se eliminan estas dos implementaciones de Greet y Bye
// func GreetAll(gs ...Greeter) {
// 	for _, g := range gs {
// 		g.Greet()
// 		fmt.Printf("Valor:\t %v\nTipo:\t %T\n\n", g, g)
// 	}
// }

// func ByeAll(bs ...Byer) {
// 	for _, b := range bs {
// 		b.Bye()
// 		fmt.Printf("Valor:\t %v\nTipo:\t %T\n\n", b, b)
// 	}
// }

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	// declaramos una variable tipo Greeter
	// que es una interface y le asignamos
	// la struct Person
	var g Greeter = Person{"Jhon"}

	// Podemos usar el metodo Greet porque Person implemeto
	// el metodo Greet de la interface Greeter
	g.Greet()

	// declaramos una vatriable tipo Greeter
	// que es una interface y le asignamos
	// un string del tipo que creamos llmado Text
	var gr Greeter = Text("Betty")

	// Podemos usar el metodo Greet porque Text impemento
	// el metodo Greet de la interface Greeter
	gr.Greet()

	fmt.Println("________________________________________________")

	p := Person{Name: "Karime"}
	var t Text = "Thor"

	// GreetAll(p, t, g, gr)

	fmt.Println("________________________________________________")

	// ByeAll(p, t)

	All(p, t)

	// Como se implemento la interface en un metodo de la structura Person
	// Se modifica el comportamiento del metodo fmt.Println
	fmt.Println(p)

	// Funciona normalmente porque no se ha implementado la interface en tipo Text
	fmt.Println(t)

}
