package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"

	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/infra/graphQL/graph/model"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}

	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// GetOrderByID is the resolver for the getOrderById field.
func (r *queryResolver) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID: id,
	}

	output, err := r.GetOrderByIdUseCase.Execute(dto.ID)
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// GetOrders is the resolver for the getOrders field.
func (r *queryResolver) GetOrders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*model.Order

	for _, order := range output {
		orders = append(orders, &model.Order{
			ID:         order.ID,
			Price:      float64(order.Price),
			Tax:        float64(order.Tax),
			FinalPrice: float64(order.FinalPrice),
		})
	}

	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
