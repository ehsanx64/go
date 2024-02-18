package lib

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rmq struct {
	Conn    *amqp.Connection
	Chann   *amqp.Channel
	Queue   amqp.Queue
	Context context.Context
	Cancel  context.CancelFunc
}

var Rmqi Rmq

func InitRabbitmq() {
	rmq := Rmq{}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	rmq.Conn = conn
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	rmq.Chann = ch
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // Name
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	FailOnError(err, "Failed to declare a queue")
	rmq.Queue = q
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	rmq.Context = ctx
	rmq.Cancel = cancel
	defer cancel()

	Rmqi = rmq
}

func Publish(message string) {
	err := Rmqi.Chann.PublishWithContext(Rmqi.Context,
		"",              // Exchange
		Rmqi.Queue.Name, // Routing key
		false,           // Mandatory
		false,           // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	FailOnError(err, "Failed to publish a message")
}
