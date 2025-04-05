package usecase

import "github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/entity"

type GetOrderByIdUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderByIdUseCaseUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderByIdUseCase) Execute(id string) (OrderOutputDTO, error) {
	order, err := c.OrderRepository.GetByID(id)
	if err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	return dto, nil
}
