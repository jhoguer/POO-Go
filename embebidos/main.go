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

type Employee struct {
	// Los campos y metodos del struct Person se
	// se convertiran en los campos y metodos de Employee
	Person
	Salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee {
	return Employee{NewPerson(name, age), salary}
}

func (e Employee) Payroll() {
	fmt.Printf("%.0f", e.Salary*0.90)
}

func main() {
	e := NewEmployee("Jhon", 34, 4500000)
	// Al declarar una variable tipo Employee
	// podemos usar sus campos y metodos
	// en este caso podemos imprimir Name y Age
	// y usar el metodo Greet que son propios de Person
	fmt.Println(e.Name)
	fmt.Println(e.Age)
	e.Greet()
	e.Payroll()

}
