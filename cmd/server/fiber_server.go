package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adzi007/ecommerce-order-service/config"
	grpcclient "github.com/adzi007/ecommerce-order-service/internal/delivery/grpc_client"
	"github.com/adzi007/ecommerce-order-service/internal/delivery/http"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-order-service/internal/infrastructure/rabbitmq"
	"github.com/adzi007/ecommerce-order-service/internal/repository"
	"github.com/adzi007/ecommerce-order-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type fiberServer struct {
	app      *fiber.App
	db       database.Database
	conn     *grpc.ClientConn
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewFiberServer(db database.Database, conn *grpc.ClientConn) Server {

	rabbitUser := config.ENV.RABBITMQ_USER
	rabbitPass := config.ENV.RABBITMQ_PASS
	rabbitHost := config.ENV.RABBITMQ_HOST
	rabbitPort := config.ENV.RABBITMQ_PORT
	rabbitVHost := config.ENV.RABBITMQ_VHOST

	amqpURL := "amqp://" + rabbitUser + ":" + rabbitPass + "@" + rabbitHost + ":" + rabbitPort + "/" + rabbitVHost

	// rabbitMQ, err := rabbitmq.NewRabbitMQ("amqp://guest:guest@localhost:5672/ecommerce_development")
	rabbitMQ, err := rabbitmq.NewRabbitMQ(amqpURL)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	fiberApp := fiber.New()

	return &fiberServer{
		app:      fiberApp,
		db:       db,
		conn:     conn,
		rabbitMQ: rabbitMQ,
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

	// log.Fatal(s.app.Listen(":5001"))

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		if err := s.app.Listen(":5002"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for termination signal
	<-stop
	log.Println("Shutting down server...")

	// Create a timeout context for shutdown (e.g., 5 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down Fiber
	if err := s.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	// Close RabbitMQ connection
	s.Close()

	log.Println("Server shut down successfully")
}

func (s *fiberServer) Close() {
	if s.rabbitMQ != nil {
		s.rabbitMQ.Close()
	}
}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	// Initialize the Cart gRPC client
	cartGrpcClient := grpcclient.NewCartGrpcClient(s.conn)

	// repository
	orderRepo := repository.NewOrderRepo(s.db)

	// use case
	orderService := service.NewOrderServiceImpl(orderRepo, cartGrpcClient, s.rabbitMQ)

	// handler
	orderHandler := http.NewOrderHttpHandle(orderService)

	// router
	s.app.Post("/", orderHandler.InsertNewOrder)
	s.app.Put("/:orderId", orderHandler.UpdateOrderStatus)
	s.app.Get("/:userId", orderHandler.GetOrdersByCustomer)
}
