package rabbit

import (
	"chainer/config"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	DefaultExchangeName = ""
	DefaultExchangeType = "direct"
	DefaultRoutingKey   = "subscriptions"
	DefaultQeueName     = "txNotifications"
)

type messageBroker struct {
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	QeueName     string
	ctx          context.Context
	cancelCtx    context.CancelFunc
}

func NewMssageBroker(con *amqp.Connection) *messageBroker {
	var b messageBroker
	b.Connection = con
	channel, err := b.Connection.Channel()
	if err != nil {
		log.Println("faild to make a rabbit channel")
	}
	b.Channel = channel
	b.ExchangeType = config.GetRabbitMQExchangeTypeString()
	if b.ExchangeType == "" {
		b.ExchangeType = DefaultExchangeType
	}
	b.ExchangeName = config.GetRabbitMQExchangeTypeString()
	if b.ExchangeName == "" {
		b.ExchangeName = DefaultExchangeName
	}
	b.RoutingKey = config.GetRabbitMQRoutingKeyString()
	if b.RoutingKey == "" {
		b.RoutingKey = DefaultRoutingKey
	}
	b.QeueName = config.GetRabbitMQQeueNameString()
	if b.QeueName == "" {
		b.QeueName = DefaultQeueName
	}
	// err = b.Channel.ExchangeDeclare(b.ExchangeName, b.ExchangeType, true, false, false, false, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Println("faild to declare a rabbit Exchange")
	// }
	return &b

}

func (b *messageBroker) Publisher(ctx context.Context, msg []byte) error {
	b.ctx, b.cancelCtx = context.WithCancel(ctx)
	defer b.cancelCtx()
	message := b.createAmqpMessage(msg)
	err := b.Channel.Publish(b.ExchangeName, b.RoutingKey, false, false, message)
	if err != nil {
		return err
	}
	return nil
}

func (b *messageBroker) Consumer(ctx context.Context, msg chan amqp.Delivery, chErr chan error) {
	_, err := b.Channel.QueueDeclare(
		b.QeueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		chErr <- err
	}
	// err = b.Channel.QueueBind(b.QeueName, b.RoutingKey, b.ExchangeName, false, nil)
	// if err != nil {
	// 	chErr <- err
	// }
	for {
		select {
		case <-ctx.Done():
			fmt.Println("what should i do with this shit ! ")

		default:
			messageChannel, err := b.Channel.Consume(
				b.QeueName,
				"",
				false,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				chErr <- err
			}

			for message := range messageChannel {
				fmt.Println(message)
				msg <- message
			}
		}
	}
}

func (b *messageBroker) createAmqpMessage(msg []byte) amqp.Publishing {
	var message amqp.Publishing
	message.ContentType = "application/json"
	message.Body = msg
	return message
}

func RegisterConsumer[T any](exchangeName string, consumer interface{}) {
	err := json.Unmarshal(msg, T)
	if err != nil {
		fmt.Printf("faild to marshal message with this err: %s", err)
	}
	consumer(T)
}
