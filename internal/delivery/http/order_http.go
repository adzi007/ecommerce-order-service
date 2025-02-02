package http

import (
	"fmt"
	"strconv"

	"github.com/adzi007/ecommerce-order-service/internal/domain"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/model"
	"github.com/gofiber/fiber/v2"
)

type orderHttpHandler struct {
	orderService domain.OrderService
}

func NewOrderHttpHandle(orderService domain.OrderService) OrderHandler {
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

	h.orderService.CreateNewOrder(reqBody)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Success create order",
	})
}

func (h *orderHttpHandler) UpdateOrderStatus(ctx *fiber.Ctx) error {

	orderId, err := strconv.ParseUint(ctx.Params("orderId"), 10, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	reqBody := new(model.UpdateStatusOrderDto)

	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := h.orderService.UpdateOrderStatus(orderId, reqBody.Status); err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"pesan": "failed to update status order",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Success Update Status Order",
	})
}
