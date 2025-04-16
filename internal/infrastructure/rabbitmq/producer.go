package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type OrderMessage struct {
	OrderID uint64 `json:"order_id"`
	UserId  string `json:"user_id"`
	Status  string `json:"status"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (r *RabbitMQ) PublishOrderStatus(exchange, routingKey string, order OrderMessage) error {

	body, err := json.Marshal(order)

	if err != nil {
		return err
	}

	// Ensure the channel is open before publishing
	if r.Channel == nil || r.Channel.IsClosed() {
		ch, err := r.Conn.Channel()
		if err != nil {
			return err
		}
		r.Channel = ch
	}

	q, err := r.Channel.QueueDeclare(
		routingKey, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		// nil,        // arguments
		amqp091.Table{"x-queue-type": "quorum"},
	)

	failOnError(err, "Failed to declare a queue")

	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.Channel.PublishWithContext(ctx,
		exchange, // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	failOnError(err, "Failed to publish a message")

	log.Printf("Published message: %+v", order)
	return nil
}
