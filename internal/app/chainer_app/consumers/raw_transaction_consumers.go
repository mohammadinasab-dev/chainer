package consumers

import (
	"chainer/internal/app/chainer_app/handlers"
	"chainer/internal/app/chainer_app/models"
	"context"

	"github.com/streadway/amqp"
)

type RawTransactionConsumer interface {
	Consumer(ctx context.Context, msg chan amqp.Delivery, chErr chan error)
}

// type rawTransactionConsumer func ( models.Tx)

func rawTransactionConsumer(msg models.Tx) {

	var parsedTransaction = handlers.Parse(msg)
	handlers.Notify(parsedTransaction)
}
