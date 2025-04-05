package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetOrderById(id string) (*Order, error)
	GetOrders() ([]*Order, error)
}
