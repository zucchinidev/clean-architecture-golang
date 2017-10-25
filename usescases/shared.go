package usescases

import (
	"github.com/zucchinidev/clean-architecture-golang/domain"
	"fmt"
)

func addItemToOrder(order *domain.Order, user *User, item domain.Item) error {
	if domainErr := order.Add(item); domainErr != nil {
		message := "Could not add item #%i "
		message += "to order #%i (of customer #%i) "
		message += "as user #%i because a business "
		message += "rule was violated: '%s'"
		err := fmt.Errorf(message,
			item.Id,
			order.Id,
			order.Customer.Id,
			user.Id,
			domainErr.Error())
		return err
	}
	return nil
}
