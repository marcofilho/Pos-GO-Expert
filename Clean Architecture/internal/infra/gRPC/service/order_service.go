package service

import (
	"context"

	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/infra/gRPC/pb"
	"github.com/marcofilho/Pos-GO-Expert/CleanArchitecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetOrderByIdUseCase usecase.GetOrderByIdUseCase
	GetOrdersUseCase    usecase.GetOrdersUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	getOrderByIdUseCase usecase.GetOrderByIdUseCase,
	getOrdersUseCase usecase.GetOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetOrderByIdUseCase: getOrderByIdUseCase,
		GetOrdersUseCase:    getOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, in *pb.GetOrderByIdRequest) (*pb.GetOrderByIdResponse, error) {
	dto := usecase.OrderInputDTO{
		ID: in.Id,
	}

	output, err := s.GetOrderByIdUseCase.Execute(dto.ID)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrderByIdResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrders(ctx context.Context, in *pb.Blank) (*pb.GetOrdersResponse, error) {
	output, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*pb.OrderResponse

	for _, order := range output {
		orders = append(orders, &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.GetOrdersResponse{Orders: orders}, nil
}
