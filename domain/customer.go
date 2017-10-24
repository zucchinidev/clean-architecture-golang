package domain

type CustomerRepository interface {
	Store(customer Customer)
	FindById(id int) Customer
}

type Customer struct {
	Id   int
	Name string
}
