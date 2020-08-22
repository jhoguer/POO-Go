package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Pagado por Paypal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado por Efectivo")
}

type CreditCard struct {
}

func (c CreditCard) Pay() {
	fmt.Println("Pagado por Tarjeta de Credito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Digite un numero de método de pago")
	fmt.Printf("1: Paypal\n2: Efectivo\n3: Tarjeta de Credito\n------> ")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Debe digitar un método valido")
	}

	if method > 3 {
		panic("Debe digitar un método valido")
	}

	payMethod := Factory(method)

	payMethod.Pay()
}
