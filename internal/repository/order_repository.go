package repository

import (
	"errors"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	"gorm.io/gorm"
)

type OrderPostgresRepo struct {
	db database.Database
}

func NewOrderRepo(db database.Database) domain.OrderRepository {
	return &OrderPostgresRepo{db: db}
}

func (r *OrderPostgresRepo) CreateNewOrder(order model.NewOrder, orderList []model.NewOrderDetail) error {

	newOrder := &model.Order{
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}
	err := r.db.GetDb().Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&newOrder).Error; err != nil {
			return err
		}

		var newOrderDetails []model.OrderDetail

		for _, val := range orderList {
			orderItem := model.OrderDetail{
				OrderID:   newOrder.ID,
				ProductID: val.ProductID,
				Quantity:  val.Quantity,
				Price:     val.Price,
			}

			newOrderDetails = append(newOrderDetails, orderItem)
		}

		if len(newOrderDetails) > 0 {
			if err := tx.Create(&newOrderDetails).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return errors.New("failed to create new order: " + err.Error())
	}

	return nil
}

func (r *OrderPostgresRepo) UpdateStatusOrder(orderId uint64, status string) error {

	var order model.Order

	if err := r.db.GetDb().First(&order, orderId).Error; err != nil {

		logger.Error().Err(err).Msgf("order not found by id %d!", orderId)

		return err
	}

	order.Status = status

	return r.db.GetDb().Save(&order).Error

}
