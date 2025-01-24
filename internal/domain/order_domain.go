package domain

import "github.com/adzi007/ecommerce-order-service/internal/model"

type OrderRepository interface {
	CreateNewOrder(model.NewOrder, []model.NewOrderDetail) error
	// SendChat(ChatBubble, string) error
	// SetReadedChat(chatRoomId, chatBubbleId string) error
	// GetChatRoomId(chatRoomId string) (*Chat, error)
	// GetChatRooms(userId string) ([]*Chat, error)
}

type OrderService interface {
	CreateNewOrder(in *model.OrderDto) error
}
