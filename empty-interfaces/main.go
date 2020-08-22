package main

import (
	"fmt"
	"strings"
)

// exampler es una interface vacia porque no se implemeta ningun metodo
// y sirve para recibir tipos de datos desconocidos
type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("Valor: \t %v\nTipo: \t %T\n\n", i, i)

	// atreves del type asertion podemos validar que tipo de dato recibimos
	// en este caso validamos el tipo concreto que trae la interfaz vacia
	// v es el valor que se esta mandando en la interface vacia
	// ok es un booleano que es true si la interface vacia es un string
	v, ok := i.(string)
	if ok {
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	}
}

func main() {
	// var e exampler
	// e.x()
	wrapper(12)
	wrapper("Betty")
	wrapper(13.67)
	wrapper(false)
	wrapper("Jhon")
}
