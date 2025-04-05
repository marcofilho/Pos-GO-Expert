package usecase

import (
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/entity"
)

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetOrders()
	if err != nil {
		return nil, err
	}

	dto := make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		dto[i] = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
	}

	return dto, nil
}
