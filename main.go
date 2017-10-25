package main

import (
	"github.com/zucchinidev/clean-architecture-golang/infrastructure"
	"github.com/zucchinidev/clean-architecture-golang/usescases"
	"net/http"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("./order.db")
	availableRepositories := infrastructure.GetAvailableRepositories()
	handlers := make(map[string]infrastructure.DBHandler)
	handlers[availableRepositories.UserRepo] = dbHandler
	handlers[availableRepositories.CustomerRepo] = dbHandler
	handlers[availableRepositories.ItemRepo] = dbHandler
	handlers[availableRepositories.OrderRepo] = dbHandler

	orderInteractor := &usescases.OrderInteractor{
		UserRepository:  infrastructure.NewDBUserRepo(handlers),
		ItemRepository:  infrastructure.NewDBItemRepo(handlers),
		OrderRepository: infrastructure.NewDBOrderRepo(handlers),
	}

	webserviceHandler := infrastructure.WebServiceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
