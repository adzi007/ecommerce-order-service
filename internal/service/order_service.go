package service

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/config"
	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/rabbitmq"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	httpclient "github.com/adzi007/ecommerce-order-service/pkg/http_client"
	"github.com/k0kubun/pp/v3"
)

type OrderService struct {
	orderRepo domain.OrderRepository
	// cartGrpcClient grpcclient.CartGrpcClient
	cartGrpcClient domain.CartRepository
	RabbitMQ       *rabbitmq.RabbitMQ
}

func NewOrderServiceImpl(orderRepo domain.OrderRepository, cartGrpcClient domain.CartRepository, rabbitMQ *rabbitmq.RabbitMQ) domain.OrderService {
	return &OrderService{
		orderRepo:      orderRepo,
		cartGrpcClient: cartGrpcClient,
		RabbitMQ:       rabbitMQ,
	}
}

type ValidateOrderItem struct {
	ID  int `json:"id"`
	Qty int `json:"qty"`
}

func (s *OrderService) CreateNewOrder(in *model.OrderDto) error {

	pp.Println("req bpdyyy >>>> ", in)

	cartResponse, err := s.cartGrpcClient.GetCartByUserID(in.UserId)

	pp.Println("cartResponse >>> ", cartResponse)

	if err != nil {
		return fmt.Errorf("failed to fetch cart data: %w", err)
	}

	var items []ValidateOrderItem
	var totalPrice uint64
	var orderDetail []model.NewOrderDetail

	for _, val := range cartResponse {
		totalPrice += val.Price * val.Qty

		items = append(items, ValidateOrderItem{
			ID:  int(val.ProductId),
			Qty: int(val.Qty),
		})

		orderDetail = append(orderDetail, model.NewOrderDetail{
			ProductID: uint(val.ProductId),
			Quantity:  int(val.Qty),
			Price:     float64(val.Price),
		})
	}

	postPayload := map[string]interface{}{
		"productsOrderList": items,
	}

	// --- validate order
	httpClient := httpclient.NewHTTPClient()

	// url := "http://localhost:3000/products/validate-order"

	url := config.ENV.API_GATEWAY + config.ENV.PRODUCT_SERVICE_PATH + "/validate-order"

	_, err = httpClient.Post(url, postPayload, nil)

	if err != nil {
		pp.Printf("POST request failed: %v", err)
	}

	// --- end validate order

	newOrder := model.NewOrder{
		UserId:        in.UserId,
		PaymentMethod: in.PaymentMethod,
		PaymentFee:    float64(in.PaymentFee),
		TotalPrice:    float64(totalPrice),
	}

	if err := s.orderRepo.CreateNewOrder(newOrder, orderDetail); err != nil {

		pp.Println("failed save new order")
		logger.Error().Err(err).Msg("failed save new order")
	}

	// delete cart by user
	_, err = s.cartGrpcClient.DeleteCartUser(in.UserId)

	if err != nil {
		pp.Println("failed delete item by userId " + in.UserId)
		logger.Error().Err(err).Msg("failed delete item by userId " + in.UserId)

		return err
	}

	return nil
}

func (s *OrderService) UpdateOrderStatus(orderId uint64, status string) error {

	userId, errUpdateStatus := s.orderRepo.UpdateStatusOrder(orderId, status)

	pp.Println("userId >>> ", userId)

	if errUpdateStatus != nil {
		logger.Error().Err(errUpdateStatus).Msgf("Failed to update status order_id %d", orderId)
		pp.Println("Failed to update status order_id ", orderId)

		return errUpdateStatus
	}

	orderMessage := rabbitmq.OrderMessage{
		OrderID: orderId,
		UserId:  userId,
		Status:  status,
	}

	err := s.RabbitMQ.PublishOrderStatus("notification", "realtime_notif", orderMessage)

	if err != nil {
		// log.Printf("Failed to publish order status: %v", err)
		logger.Error().Err(err).Msgf("Failed to publish order status: %v", err)
		return err
	}

	return nil
}

func (s *OrderService) GetOrderByUser(userId string) ([]model.Order, error) {

	orders, err := s.orderRepo.GetOrderByUser(userId)

	if err != nil {
		pp.Println("Error fetching orders:", err)
		return nil, err
	}

	return orders, nil

}
