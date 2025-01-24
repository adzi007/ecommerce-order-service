package service

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	"github.com/adzi007/ecommerce-order-service/internal/repository"
)

type OrderService struct {
	orderRepo repository.OrderPostgresRepo
}

func NewOrderServiceImpl(orderRepo repository.OrderPostgresRepo) domain.OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateNewOrder(in *model.OrderDto) error {

	fmt.Println("body request >>> ", in)

	// if err:= s.orderRepo.CreateNewOrder()

	return nil
}
