package main

import (
	"time"

	"github.com/adzi007/ecommerce-order-service/cmd/server"
	"github.com/adzi007/ecommerce-order-service/config"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	grpcconnection "github.com/adzi007/ecommerce-order-service/internal/infrastructure/grpc_connection"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/monitoring"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewPostgreesDatabase()

	// conn := grpcconnection.NewGrpcConnection("localhost:9001")

	grpcPort := config.ENV.GRPC_PORT

	conn := grpcconnection.NewGrpcConnection("ecommerce-cart-service:" + grpcPort)
	defer conn.Close()

	servernya := server.NewFiberServer(db, conn)

	servernya.Use(limiter.New(limiter.Config{
		Max:        10,               // 10 requests
		Expiration: 30 * time.Second, // per 30 seconds
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // limit per IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	// Register Prometheus metrics
	monitoring.RegisterMetrics()

	servernya.Use(monitoring.PrometheusMiddleware())

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	servernya.Start()
}
