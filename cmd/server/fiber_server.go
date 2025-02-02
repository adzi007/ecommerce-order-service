package server

import (
	"log"

	grpcclient "github.com/adzi007/ecommerce-order-service/internal/delivery/grpc_client"
	"github.com/adzi007/ecommerce-order-service/internal/delivery/http"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/repository"
	"github.com/adzi007/ecommerce-order-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conn *grpc.ClientConn
}

func NewFiberServer(db database.Database, conn *grpc.ClientConn) Server {

	fiberApp := fiber.New()

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conn: conn,
	}
}

func (s *fiberServer) Use(args interface{}) {
	s.app.Use(args)
}

func (s *fiberServer) Start() {
	// Define routes

	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello from Fiber! ini pesan dari admin")
	})

	logger.Info().Msg("This is an info message")

	logger.Warn().Msg("This is a warning message")

	s.initializeCartServiceHttpHandler()

	log.Fatal(s.app.Listen(":5001"))
}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	// Initialize the Cart gRPC client
	cartGrpcClient := grpcclient.NewCartGrpcClient(s.conn)

	// repository
	orderRepo := repository.NewOrderRepo(s.db)

	// use case
	orderService := service.NewOrderServiceImpl(orderRepo, cartGrpcClient)

	// handler
	orderHandler := http.NewOrderHttpHandle(orderService)

	// router
	s.app.Post("/", orderHandler.InsertNewOrder)
	s.app.Put("/:orderId", orderHandler.UpdateOrderStatus)

}
