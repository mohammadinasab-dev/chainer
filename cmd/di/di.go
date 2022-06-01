package di

import (
	"chainer/internal/app/chainer_app/controllers"
	"chainer/internal/app/chainer_app/handlers"
	repo "chainer/internal/app/chainer_app/repositories"
	postgres "chainer/internal/pkg/postgres"
	"chainer/internal/pkg/rabbit"

	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
)

var (
	dbRepo repo.DbReposiroty
)

func InitDB() *pg.DB {
	return provideDB()
}

func InitRabbitMq() *amqp.Connection {
	return provideRbbitMq()
}

func NewDBDataSource() repo.DbReposiroty {
	if dbRepo == nil {
		dbRepo = postgres.NewDBDataSource(provideDB())
	}
	return dbRepo
}

func AddressGenerator() *controllers.Address {
	pgCon := provideDB()
	addressDs := postgres.NewAddressRepository(pgCon)
	publisher := rabbit.NewMssageBroker(provideRbbitMq())
	addressHandler := handlers.NewAddress(addressDs, publisher)
	address := controllers.NewAddress(addressHandler)
	return address
}

func RawTransactionHandler() *handlers.RawTransaction {
	msg := make(chan amqp.Delivery)
	chErr := make(chan error)
	consumer := rabbit.NewMssageBroker(provideRbbitMq())
	rawTransaction := handlers.NewRawTransaction(consumer, msg, chErr)
	return rawTransaction
}
