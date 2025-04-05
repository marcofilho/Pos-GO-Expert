package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetByID(id string) (*Order, error)
	GetOrders() ([]*Order, error)
}
