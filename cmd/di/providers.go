package di

import (
	"chainer/config"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
)

var (
	pgCon       *pg.DB
	rabbitMqCon *amqp.Connection
)

func provideDB() *pg.DB {
	if pgCon != nil {
		return pgCon
	}
	options := &pg.Options{
		User:     config.GetDatabaseConnectionUser(),
		Password: config.GetDatabaseConnectionPass(),
		Addr:     config.GetDatabaseConnectionUrl(),
		Database: config.GetDatabaseConnectionDB(),
		PoolSize: 10,
	}
	pgCon = pg.Connect(options)
	if pgCon == nil {
		log.Panic("faild to connect to the database")
	}

	return pgCon
}

func provideRbbitMq() *amqp.Connection {
	url := config.GetRabbitMQConnectionString()
	rabbitMqCon, err := amqp.Dial(url)
	if err != nil {
		log.Panic("faild to connect to the rabbitMq")
	}
	return rabbitMqCon
}
