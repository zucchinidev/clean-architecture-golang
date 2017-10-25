package infrastructure

import (
	"github.com/zucchinidev/clean-architecture-golang/usescases"
	"fmt"
)

func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
	availableRepositories := GetAvailableRepositories()
	dbUserRepo := new(DBUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers[availableRepositories.UserRepo]
	return dbUserRepo
}

func (repo *DBUserRepo) Store(user usescases.User) {
	isAdmin := "no"
	if user.IsAdmin {
		isAdmin = "yes"
	}
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, customer_id, is_admin)
												VALUES ('%d', '%d', '%v')`, user.Id, user.Customer.Id, isAdmin))
	customerRepo := NewDBCustomerRepo(repo.dbHandlers)
	customerRepo.Store(user.Customer)
}

func (repo *DBUserRepo) FindById(userId int) usescases.User {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT is_admin, customer_id
													FROM users WHERE id = '%d' LIMIT 1`, userId))
	var isAdmin string
	var customerId int
	row.Next()
	row.Scan(&isAdmin, &customerId)
	customerRepo := NewDBCustomerRepo(repo.dbHandlers)
	customer := customerRepo.FindById(customerId)
	u := usescases.User{
		Id: userId,
		Customer: customer,
	}
	u.IsAdmin = false
	if isAdmin == "yes" {
		u.IsAdmin = true
	}
	return u
}
