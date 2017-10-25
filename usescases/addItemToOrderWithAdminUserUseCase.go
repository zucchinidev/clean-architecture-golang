package usescases

import (
	"fmt"
	"github.com/zucchinidev/clean-architecture-golang/domain"
)

type AdminOrderInteractor struct {
	OrderInteractor
}

func (interactor *AdminOrderInteractor) Add(userId, orderId, itemId int) error {
	user := interactor.UserRepository.FindById(userId)
	order := interactor.OrderRepository.FindById(orderId)

	if err := userIsAdmin(&user, &order); err != nil {
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
		"Admin added item '%s' (#%i) to order #%i",
		item.Name, item.Id, order.Id))
	return nil
}

func userIsAdmin(user *User, order *domain.Order) error {
	if !user.IsAdmin {
		message := "User #%i (customer #%i) "
		message += "is not allowed to add items "
		message += "to order #%i (of customer #%i), "
		message += "because he is not an administrator"
		err := fmt.Errorf(message,
			user.Id,
			user.Customer.Id,
			order.Id,
			order.Customer.Id)
		return err
	}
	return nil
}
