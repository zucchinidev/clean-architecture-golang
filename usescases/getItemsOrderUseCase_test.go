package usescases

import (
	"github.com/zucchinidev/clean-architecture-golang/domain"
	"fmt"
)

func (interactor *OrderInteractor) Items(userId, orderId int) ([]Item, error) {
	var items []Item
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)
	if err := checkIsAllowedSeeItems(&user, &order); err != nil {
		interactor.Logger.Log(err.Error())
		return items, err
	}

	items = make([]Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = Item{item.Id, item.Name, item.Value}
	}
	return items, nil
}

func checkIsAllowedSeeItems(user *User, order *domain.Order) error {
	if user.Customer.Id != order.Customer.Id {
		message := "User #%i (customer #%i) "
		message += "is not allowed to see items "
		message += "in order #%i (of course #%i) "
		return fmt.Errorf(message,
			user.Id,
			user.Customer.Id,
			order.Id,
			order.Customer.Id)
	}
	return nil
}
