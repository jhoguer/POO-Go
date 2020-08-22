package main

import "fmt"

// exampler es una interface vacia porque no se implemeta ningun metodo
// y sirve para recibir tipos de datos desconocidos
type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("Valor: \t %v\nTipo: \t %T\n\n", i, i)
}

func main() {
	// var e exampler
	// e.x()
	wrapper(12)
	wrapper(13.67)
	wrapper(false)
	wrapper("Jhon")
}
