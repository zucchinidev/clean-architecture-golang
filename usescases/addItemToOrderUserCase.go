package usescases

import (
	"fmt"
	"github.com/zucchinidev/clean-architecture-golang/domain"
)

func (interactor *OrderInteractor) Add(userId, orderId, itemId int) error {
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)
	if err := checkIsAllowedAddItems(&user, &order); err != nil {
		interactor.Logger.Log(err.Error())
		return err
	}
	item := interactor.ItemRepository.FindById(itemId)
	if err := addItemToOrder(&order, &user, item); err != nil {
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.OrderRepository.Store(order)
	interactor.Logger.Log(fmt.Sprintf(
		"User added item '%s' (#%i) to order #%i",
		item.Name, item.Id, order.Id))
	return nil
}

func checkIsAllowedAddItems(user *User, order *domain.Order) error {
	if user.Customer.Id != order.Customer.Id {
		message := "User #%i (customer #%i) "
		message += "is not allowed to add items "
		message += "to order #%i (of customer #%i)"
		return fmt.Errorf(message,
			user.Id,
			user.Customer.Id,
			order.Id,
			order.Customer.Id)
	}
	return nil
}
