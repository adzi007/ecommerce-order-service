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
	// conf *config.Config
}

func NewFiberServer(db database.Database, conn *grpc.ClientConn) Server {
	fiberApp := fiber.New()
	// fiberApp.Logger.SetLevel(log.DEBUG)

	// fiberApp.Get("/docs/*", swagger.HandlerDefault)

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
	// logger.Warn().Str("user", "john_doe").Msg("This is a warning message")
	logger.Warn().Msg("This is a warning message")

	s.initializeCartServiceHttpHandler()

	log.Fatal(s.app.Listen(":5001"))
}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	// ctx := context.Background()

	// redisRepo := cachestore.NewRedisCache(ctx, "localhost:6379", "", 0)

	// ------ gRpc Setup -------------------

	// conn := grpcconnection.NewGrpcConnection("localhost:9001")
	// defer conn.Close()

	// // Initialize the Cart gRPC client
	cartGrpcClient := grpcclient.NewCartGrpcClient(s.conn)

	// -------------------------------------

	// repository
	orderRepo := repository.NewOrderRepo(s.db)

	// product service repository
	// productServiceRepo := productservicerepo.NewProductServiceRepository()

	// use case
	orderService := service.NewOrderServiceImpl(orderRepo, cartGrpcClient)

	// handler

	orderHandler := http.NewOrderHttpHandle(orderService)

	// router
	// s.app.Post("/cart", cartHandler.InsertNewCart)
	s.app.Post("/", orderHandler.InsertNewOrder)
	// s.app.Get("/", cartHandler.GetCustomerCart)
	// s.app.Get("/:userId", cartHandler.GetCartByCustomer)
	// s.app.Put("/", cartHandler.UpdateQty)
	// s.app.Delete("/:cartId", cartHandler.DeleteCartItem)
	// s.app.Get("/check/redis", cartHandler.Check)

}
