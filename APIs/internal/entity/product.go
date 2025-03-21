package entity

import (
	"errors"
	"time"

	"github.com/marcofilho/Pos-GO-Expert/APIs/pkg/entity"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID    = errors.New("invalid id")
	ErrNameRequired = errors.New("name is required")
	ErrPriceInvalid = errors.New("price is invalid")
	ErrInvalidPrice = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(string(p.ID.String())); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameRequired
	}

	if p.Price <= 0 {
		return ErrPriceInvalid
	}

	return nil
}
