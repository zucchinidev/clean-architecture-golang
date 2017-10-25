package usescases

import "github.com/zucchinidev/clean-architecture-golang/domain"

const checkMark = "\u2713"
const ballotX = "\u2717"
const customerId = 1

type FakeUserRepository struct {
}

func (f *FakeUserRepository) FindById(userId int) User {
	return User{
		Id: userId,
		Customer: struct {
			Id   int
			Name string
		}{
			Id: customerId, Name: "fake customer",
		},
	}
}

func (f *FakeUserRepository) Store(user User) {

}

type FakeOrderRepository struct {
}

func (f *FakeOrderRepository) FindById(orderId int) domain.Order {
	items := []domain.Item{
		{
			Id:        1,
			Value:     1,
			Name:      "fake item 1",
			Available: true,
		},
		{
			Id:        2,
			Value:     2,
			Name:      "fake item 2",
			Available: true,
		},
	}
	return domain.Order{
		Id:    orderId,
		Items: items,
		Customer: struct {
			Id   int
			Name string
		}{
			Id: customerId, Name: "fake customer",
		},
	}
}

func (f *FakeOrderRepository) Store(order domain.Order) {

}

type FakeItemRepository struct {
}

func (f *FakeItemRepository) FindById(itemId int) domain.Item {
	return domain.Item{
		Id: itemId,
		Available: true,
	}
}

func (f *FakeItemRepository) Store(item domain.Item) {

}

type FakeLogger struct {
}

func (f *FakeLogger) Log(message string) error {
	return nil
}

func getFakeOrderInteractor() OrderInteractor {
	return OrderInteractor{
		UserRepository:  new(FakeUserRepository),
		OrderRepository: new(FakeOrderRepository),
		ItemRepository:  new(FakeItemRepository),
		Logger:          new(FakeLogger),
	}
}
