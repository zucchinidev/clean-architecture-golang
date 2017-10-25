package infrastructure

import (
	"github.com/zucchinidev/clean-architecture-golang/domain"
	"fmt"
)

func NewDBOrderRepo(dbHandlers map[string]DBHandler) *DBOrderRepo {
	dbOrderRepo := new(DBOrderRepo)
	dbOrderRepo.dbHandlers = dbHandlers
	dbOrderRepo.dbHandler = dbHandlers[GetAvailableRepositories().OrderRepo]
	return dbOrderRepo
}

func (repo *DBOrderRepo) Store(order domain.Order) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO orders (id, customer_id)
                                        VALUES ('%d', '%v')`,
		order.Id, order.Customer.Id))
	for _, item := range order.Items {
		repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items2orders (item_id, order_id)
                                            VALUES ('%d', '%d')`,
			item.Id, order.Id))
	}
}

func (repo *DBOrderRepo) FindById(id int) domain.Order {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT customer_id FROM orders
                                             WHERE id = '%d' LIMIT 1`,
		id))
	var customerId int
	row.Next()
	row.Scan(&customerId)
	customerRepo := NewDBCustomerRepo(repo.dbHandlers)
	order := domain.Order{Id: id, Customer: customerRepo.FindById(customerId)}
	var itemId int
	itemRepo := NewDBItemRepo(repo.dbHandlers)
	row = repo.dbHandler.Query(fmt.Sprintf(`SELECT item_id FROM items2orders
                                            WHERE order_id = '%d'`,
		order.Id))
	for row.Next() {
		row.Scan(&itemId)
		order.Add(itemRepo.FindById(itemId))
	}
	return order
}
