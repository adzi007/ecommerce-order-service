package service

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/model"
)

type OrderService struct {
	orderRepo domain.OrderRepository
	// cartGrpcClient grpcclient.CartGrpcClient
	cartGrpcClient domain.CartRepository
}

func NewOrderServiceImpl(orderRepo domain.OrderRepository, cartGrpcClient domain.CartRepository) domain.OrderService {
	return &OrderService{
		orderRepo:      orderRepo,
		cartGrpcClient: cartGrpcClient,
	}
}

func (s *OrderService) CreateNewOrder(in *model.OrderDto) error {

	// fmt.Println("body request >>> ", in)

	cartResponse, err := s.cartGrpcClient.GetCartByUserID(in.UserId)

	if err != nil {
		return fmt.Errorf("failed to fetch cart data: %w", err)
	}

	fmt.Println("cartResponse >>> ", cartResponse)

	// if err:= s.orderRepo.CreateNewOrder()

	return nil
}
