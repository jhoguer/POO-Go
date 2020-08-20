package customer

// Customer es the struct of a client
type Customer struct {
	name    string
	address string
	phone   string
}

// New returns a new Customer
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
