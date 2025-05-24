package service

import (
	"errors"
	"testing"

	"github.com/adzi007/ecommerce-order-service/internal/model"
	"github.com/stretchr/testify/assert"
)

// ---- mock order repo ----

type mockOrderRepo struct {
	mockGetOrderByUser func(userId string) ([]model.Order, error)
}

func (m *mockOrderRepo) CreateNewOrder(o model.NewOrder, d []model.NewOrderDetail) error {
	return nil
}
func (m *mockOrderRepo) UpdateStatusOrder(orderId uint64, status string) (string, error) {
	return "", nil
}
func (m *mockOrderRepo) GetOrderByUser(userId string) ([]model.Order, error) {
	return m.mockGetOrderByUser(userId)
}

func TestGetOrderByUser_Success(t *testing.T) {
	mockOrders := []model.Order{
		{ID: 1, UserId: "user123", TotalPrice: 100.0},
		{ID: 2, UserId: "user123", TotalPrice: 200.0},
	}

	mockRepo := &mockOrderRepo{
		mockGetOrderByUser: func(userId string) ([]model.Order, error) {
			return mockOrders, nil
		},
	}

	orderService := NewOrderServiceImpl(mockRepo, nil, nil)

	orders, err := orderService.GetOrderByUser("user123")

	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, mockOrders, orders)
}

func TestGetOrderByUser_Error(t *testing.T) {

	mockRepo := &mockOrderRepo{
		mockGetOrderByUser: func(userId string) ([]model.Order, error) {
			return nil, errors.New("DB error")
		},
	}

	orderService := NewOrderServiceImpl(mockRepo, nil, nil)

	orders, err := orderService.GetOrderByUser("user123")

	assert.Error(t, err)
	assert.Nil(t, orders)
}
