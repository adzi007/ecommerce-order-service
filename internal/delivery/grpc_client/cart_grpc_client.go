package grpcclient

import (
	"context"

	proto "github.com/adzi007/ecommerce-order-service/cart_proto"
	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"google.golang.org/grpc"
)

type CartGrpcClient struct {
	// client cart_proto.CartServiceClient
	client proto.CartServiceClient
}

func NewCartGrpcClient(conn *grpc.ClientConn) domain.CartRepository {
	return &CartGrpcClient{
		// client: cart_proto.NewCartServiceClient(conn),
		client: proto.NewCartServiceClient(conn),
	}
}

func (c *CartGrpcClient) GetCartByUserID(userID string) ([]domain.CartItem, error) {
	req := &proto.CartRequest{
		Id: userID,
	}

	resp, err := c.client.GetCartUser(context.Background(), req)
	if err != nil {
		return nil, err
	}

	// Map response to domain.CartItem
	var cartItems []domain.CartItem
	for _, item := range resp.Data {
		cartItems = append(cartItems, domain.CartItem{
			ID:    item.Id,
			Name:  item.Name,
			Slug:  item.Slug,
			Price: item.Price,
			Qty:   item.Qty,
			Category: domain.ProductCategory{
				Name: item.Category.Name,
				Slug: item.Category.Slug,
			},
		})
	}

	return cartItems, nil
}
