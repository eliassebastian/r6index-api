package rabbitmq

import (
	"context"
	"log"

	"github.com/eliassebastian/r6index-api/pkg/auth"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	auth       *auth.AuthStore
}

func New(auth *auth.AuthStore) (*RabbitConsumer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"r6index", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,    // queue name
		"",        // routing key
		"r6index", // exchange
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &RabbitConsumer{
		connection: conn,
		channel:    ch,
		queue:      &q,
		auth:       auth,
	}, nil
}

func (r *RabbitConsumer) Consume(ctx context.Context) {
	log.Println("RabbitMQ Consumer Running")

	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Exiting RabbitMQ Loop")
			return
		case msg := <-msgs:
			log.Println("New Message")
			err := r.auth.Write(msg.Body)
			if err != nil {
				log.Println("could not write to session cache")
			}
		}
	}
}

func (r *RabbitConsumer) Close() {
	log.Println("closing rabbitmq connection")
	r.channel.Close()
	r.connection.Close()
}
