package http

import "github.com/gofiber/fiber/v2"

type OrderHandler interface {
	InsertNewOrder(ctx *fiber.Ctx) error
	UpdateOrderStatus(ctx *fiber.Ctx) error
}
