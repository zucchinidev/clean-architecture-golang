package infrastructure

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DBHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type DbRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

type DBUserRepo DbRepo
type DBCustomerRepo DbRepo
type DBOrderRepo DbRepo
type DBItemRepo DbRepo

type AvailableRepositories struct {
	UserRepo     string
	CustomerRepo string
	ItemRepo     string
	OrderRepo    string
}

func GetAvailableRepositories() AvailableRepositories {
	return AvailableRepositories{
		UserRepo:     "UserRepo",
		CustomerRepo: "CustomerRepo",
		ItemRepo:     "ItemRepo",
		OrderRepo:    "OrderRepo",
	}
}
