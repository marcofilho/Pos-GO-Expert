package graph

import "github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetOrderByIdUseCase usecase.GetOrderByIdUseCase
	GetOrdersUseCase    usecase.GetOrdersUseCase
}
