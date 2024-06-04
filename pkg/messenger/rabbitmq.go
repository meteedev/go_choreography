package messenger

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	Ctx  context.Context
}

func NewRabbitMQ(ctx context.Context, amqpURL string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		Conn: conn,
		Ch:   ch,
		Ctx:  ctx,
	}, nil
}

func (r *RabbitMQ) DeclareExchange(name, kind string, durable bool) error {
	return r.Ch.ExchangeDeclare(
		name,    // name
		kind,    // type
		durable, // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
}

func (r *RabbitMQ) DeclareQueue(name string, durable bool) (amqp.Queue, error) {
	return r.Ch.QueueDeclare(
		name,    // name
		durable, // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
}

func (r *RabbitMQ) BindQueue(queueName, exchangeName, routingKey string) error {
	return r.Ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

// Close closes the connection and channel.
func (r *RabbitMQ) Close() error {
	var err error
	if r.Ch != nil {
		if e := r.Ch.Close(); e != nil {
			err = e
		}
	}
	if r.Conn != nil {
		if e := r.Conn.Close(); e != nil {
			err = e
		}
	}
	return err
}

func (r *RabbitMQ) GracefulShutdown(ctx context.Context) {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stopChan:
		log.Println("RabbitMQ Received shutdown signal")
		if err := r.Close(); err != nil {
			log.Println("Failed to close RabbitMQ connection:", err)
		}
	case <-ctx.Done():
		log.Println("Context done:", ctx.Err())
	}
}

func (r *RabbitMQ) Publish(ctx context.Context, msg []byte, destinationName string, isExchange bool) error {
	var exchangeName, routingKey string

	if isExchange {
		exchangeName = destinationName
	} else {
		routingKey = destinationName
	}

	return r.Ch.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}

func (r *RabbitMQ) ConsumeMessage(ctx context.Context, queueName string) (<-chan []byte, error) {

	msgs, err := r.Ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	out := make(chan []byte)
	go func() {
		defer close(out)
		for {
			select {
			case d, ok := <-msgs:
				if !ok {
					log.Println("RabbitMQ message channel closed.")
					return
				}

				var msg []byte
				msg = d.Body
				// if err := json.Unmarshal(d.Body, &msg); err != nil {
				// 	log.Printf("Error un marshalling message: %v", err)
				// 	continue
				// }
				out <- msg
			case <-ctx.Done():
				log.Println("Context done, stopping consumer.")
				return
			}
		}
	}()
	return out, nil
}

func (r *RabbitMQ) Consume(ctx context.Context, queueName string, handlerFunc func(ctx context.Context, payload []byte) error) error {

	msgs, err := r.ConsumeMessage(ctx, queueName)
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-msgs:
			//log.Println("Consumed message:", msg)
			handlerFunc(ctx, msg)
		case <-ctx.Done():
			return nil
		}
	}
}
