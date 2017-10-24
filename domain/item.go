package domain

type ItemRepository interface {
	Store(item Item)
	FindById(id int) Item
}

type Item struct {
	Id        int
	Name      string
	Value     float64
	Available bool
}
