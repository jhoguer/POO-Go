package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

type Employee struct {
	// Los campos y metodos del struct Person se
	// se convertiran en los campos y metodos de Employee
	Person
	Human  // Human tine el campo Age igual que Person, esto genera ambiguedad
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

func (e Employee) Payroll() {
	fmt.Printf("%.0f", e.Salary*0.90)
}

func (e Employee) Greet() {
	fmt.Println("Saludo desde Empleado")
}

func main() {
	e := NewEmployee("Jhon", 34, 3, 4500000)
	// Al declarar una variable tipo Employee
	// podemos usar sus campos y metodos
	// en este caso podemos imprimir Name y Age
	// y usar el metodo Greet que son propios de Person
	fmt.Println(e.Name)

	// Para resolver la ambiguedad se debe indicar cual edad es la que vamos a usar
	// si Age de Human o Age de Person
	fmt.Println(e.Human.Age)
	fmt.Println(e.Person.Age)

	// usamos en metodo Greet de Person
	e.Person.Greet()

	// usamos el metodo Greet de Empleado
	e.Greet()
	e.Payroll()

}
