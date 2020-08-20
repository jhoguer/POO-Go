package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// declaracion de alias
type myAlias = course

// definicion de tipo
// newCourse tiene los campos del tipo base, que es course
// en este caso tiene el campo name, pero no hereda sus metodos
// y el tipo de la variable c del ejemplo cambia a tipo newCourse
type newCourse course

// tipo newBool creado a partir del tipo bool
type newBool bool

// funcion asignada al nuevo tipo creado newBool
func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	c := newCourse{name: "Go"}
	//c.Print()
	fmt.Printf("El tipo es: %T\n", c)

	// se declara una variable tipo newBool
	var b newBool = true

	// se usa el metodo que creamos para el nuevo tipo newBool
	fmt.Println(b.String())
}
