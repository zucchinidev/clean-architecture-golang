package usescases

import "github.com/zucchinidev/clean-architecture-golang/domain"

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}

type User struct {
	Id       int
	IsAdmin  bool
	Customer domain.Customer
}

type Item struct {
	Id    int
	Name  string
	Value float64
}

type Logger interface {
	Log(message string) error
}

type OrderInteractor struct {
	UserRepository  UserRepository
	OrderRepository domain.OrderRepository
	ItemRepository  domain.ItemRepository
	Logger          Logger
}
