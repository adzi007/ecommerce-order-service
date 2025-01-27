package service

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	httpclient "github.com/adzi007/ecommerce-order-service/pkg/http_client"
	"github.com/k0kubun/pp/v3"
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

type ValidateOrderItem struct {
	ID  int `json:"id"`
	Qty int `json:"qty"`
}

func (s *OrderService) CreateNewOrder(in *model.OrderDto) error {

	cartResponse, err := s.cartGrpcClient.GetCartByUserID(in.UserId)

	if err != nil {
		return fmt.Errorf("failed to fetch cart data: %w", err)
	}

	var items []ValidateOrderItem

	for _, val := range cartResponse {
		items = append(items, ValidateOrderItem{
			ID:  int(val.ProductId),
			Qty: int(val.Qty),
		})
	}

	postPayload := map[string]interface{}{
		"productsOrderList": items,
	}

	httpClient := httpclient.NewHTTPClient()
	url := "http://localhost:3000/products/validate-order"

	response, err := httpClient.Post(url, postPayload, nil)

	if err != nil {
		pp.Printf("POST request failed: %v", err)
	}

	pp.Printf("POST Response: %s\n", string(response))

	return nil
}
