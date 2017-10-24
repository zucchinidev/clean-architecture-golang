package domain

import "errors"

type Order struct {
	Id       int
	Customer Customer
	Items    []Item
}

var ErrorItemNotAvailable = errors.New("cannot add unavailable items to order")
var ErrorOrderExceedsTheMaxValue = errors.New("an order may not exceed a total value of $250")

func (order *Order) Add(item Item) error {
	if !item.Available {
		return ErrorItemNotAvailable
	}
	if order.value()+item.Value > 250.00 {
		return ErrorOrderExceedsTheMaxValue
	}
	order.Items = append(order.Items, item)
	return nil
}

func (order *Order) value() float64 {
	sum := 0.0
	for item := range order.Items {
		sum = sum + order.Items[item].Value
	}
	return sum
}

type OrderRepository interface {
	Store(item Order)
	FindById(id int) Order
}
