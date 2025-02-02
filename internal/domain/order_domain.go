package domain

import "github.com/adzi007/ecommerce-order-service/internal/model"

type OrderRepository interface {
	CreateNewOrder(model.NewOrder, []model.NewOrderDetail) error
	UpdateStatusOrder(orderId uint64, status string) error
	// SendChat(ChatBubble, string) error
	// SetReadedChat(chatRoomId, chatBubbleId string) error
	// GetChatRoomId(chatRoomId string) (*Chat, error)
	// GetChatRooms(userId string) ([]*Chat, error)
}

type OrderService interface {
	CreateNewOrder(in *model.OrderDto) error
	UpdateOrderStatus(orderId uint64, status string) error
}
