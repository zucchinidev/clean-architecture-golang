package infrastructure

import (
	"github.com/zucchinidev/clean-architecture-golang/domain"
	"fmt"
)

func NewDBItemRepo(dbHandlers map[string]DBHandler) *DBItemRepo {
	dbItemRepo := new(DBItemRepo)
	dbItemRepo.dbHandlers = dbHandlers
	dbItemRepo.dbHandler = dbHandlers["DBItemRepo"]
	return dbItemRepo
}

func (repo *DBItemRepo) Store(item domain.Item) {
	available := "no"
	if item.Available {
		available = "yes"
	}
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items (id, name, value, available)
                                        VALUES ('%d', '%v', '%f', '%v')`,
		item.Id, item.Name, item.Value, available))
}

func (repo *DBItemRepo) FindById(id int) domain.Item {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, value, available
                                             FROM items WHERE id = '%d' LIMIT 1`,
		id))
	var name string
	var value float64
	var available string
	row.Next()
	row.Scan(&name, &value, &available)
	item := domain.Item{Id: id, Name: name, Value: value}
	item.Available = false
	if available == "yes" {
		item.Available = true
	}
	return item
}
