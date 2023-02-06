package rabbitmq

import (
	"context"
	"log"

	"github.com/eliassebastian/r6index-api/pkg/auth"
	"github.com/eliassebastian/r6index-api/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	auth       *auth.AuthStore
}

func New(auth *auth.AuthStore) (*RabbitConsumer, error) {
	url := utils.GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitConsumer{
		connection: conn,
		channel:    ch,
		auth:       auth,
	}, nil
}

func (r *RabbitConsumer) Consume(ctx context.Context) {
	msgs, err := r.channel.Consume(
		"auth", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgs:
			err := r.auth.Write(msg.Body)
			if err != nil {
				log.Println("could not write to session cache")
			}
		}
	}
}

func (r *RabbitConsumer) Close() {
	r.channel.Close()
	r.connection.Close()
}
