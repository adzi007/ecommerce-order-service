package main

import (
	"github.com/adzi007/ecommerce-order-service/cmd/server"
	"github.com/adzi007/ecommerce-order-service/config"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/gofiber/contrib/fiberzerolog"
)

func main() {
	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewPostgreesDatabase()

	servernya := server.NewFiberServer(db)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	servernya.Start()
}
