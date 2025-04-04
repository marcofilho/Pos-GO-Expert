package database

import (
	"database/sql"

	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.DB.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetOrderById(id string) (*entity.Order, error) {
	row := o.DB.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", id)

	var order entity.Order

	if err := row.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
		return nil, err
	}
	return &order, nil

}

func (o *OrderRepository) GetOrders() ([]*entity.Order, error) {
	rows, err := o.DB.Query("SELECT id, price, tax, final_price FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entity.Order

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}
