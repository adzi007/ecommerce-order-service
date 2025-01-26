package main

import (
	"fmt"

	"github.com/adzi007/ecommerce-order-service/cmd/server"
	"github.com/adzi007/ecommerce-order-service/config"
	grpcclient "github.com/adzi007/ecommerce-order-service/internal/delivery/grpc_client"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	grpcconnection "github.com/adzi007/ecommerce-order-service/internal/infrastructure/grpc_connection"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/service"
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

	conn := grpcconnection.NewGrpcConnection("localhost:9001")
	defer conn.Close()

	// 2. Initialize gRPC client
	cartGrpcClient := grpcclient.NewCartGrpcClient(conn)

	// 3. Create Usecase
	cartUsecase := service.NewCartUsecase(cartGrpcClient)

	cartItems, err := cartUsecase.GetCartByUserID("4973b3ac-c8a9-416f-b145-8b14422ce848")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("cartItems >>> ", cartItems)

	servernya.Start()
}
