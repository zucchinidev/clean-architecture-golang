package infrastructure

import (
	"github.com/zucchinidev/clean-architecture-golang/domain"
	"fmt"
)

func NewDBCustomerRepo(dbHandlers map[string]DBHandler) *DBCustomerRepo {
	dbCustomerRepo := new(DBCustomerRepo)
	dbCustomerRepo.dbHandlers = dbHandlers
	dbCustomerRepo.dbHandler = dbHandlers[GetAvailableRepositories().CustomerRepo]
	return dbCustomerRepo
}

func (repo *DBCustomerRepo) Store(customer domain.Customer) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO customers (id, name)
                                        VALUES ('%d', '%v')`,
		customer.Id, customer.Name))
}

func (repo *DBCustomerRepo) FindById(id int) domain.Customer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name FROM customers
                                             WHERE id = '%d' LIMIT 1`,
		id))
	var name string
	row.Next()
	row.Scan(&name)
	return domain.Customer{Id: id, Name: name}
}

