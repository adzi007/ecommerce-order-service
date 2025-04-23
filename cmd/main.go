package main

import (
	"github.com/adzi007/ecommerce-order-service/cmd/server"
	"github.com/adzi007/ecommerce-order-service/config"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	grpcconnection "github.com/adzi007/ecommerce-order-service/internal/infrastructure/grpc_connection"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/monitoring"
	"github.com/gofiber/contrib/fiberzerolog"
)

func main() {
	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewPostgreesDatabase()

	// conn := grpcconnection.NewGrpcConnection("localhost:9001")
	conn := grpcconnection.NewGrpcConnection("ecommerce-cart-service:9001")
	defer conn.Close()

	servernya := server.NewFiberServer(db, conn)

	// Register Prometheus metrics
	monitoring.RegisterMetrics()

	servernya.Use(monitoring.PrometheusMiddleware())

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	servernya.Start()
}
