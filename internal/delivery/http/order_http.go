package http

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	"github.com/adzi007/ecommerce-order-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type orderHttpHandler struct {
	orderService service.OrderService
}

func NewOrderHttpHandle(orderService service.OrderService) OrderHandler {
	return &orderHttpHandler{
		orderService: orderService,
	}
}

func (h *orderHttpHandler) InsertNewOrder(ctx *fiber.Ctx) error {

	reqBody := new(model.OrderDto)

	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	fmt.Println("reqBody >>> ", reqBody)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "success create a new cart 1",
	})
}
