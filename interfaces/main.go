package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("Valor:\t %v\nTipo:\t %T\n\n", g, g)
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

	GreetAll(p, t, g, gr)

}
