package service

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/model"
)

type OrderService struct {
	orderRepo domain.OrderRepository
}

func NewOrderServiceImpl(orderRepo domain.OrderRepository) domain.OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateNewOrder(in *model.OrderDto) error {

	fmt.Println("body request >>> ", in)

	// if err:= s.orderRepo.CreateNewOrder()

	return nil
}
