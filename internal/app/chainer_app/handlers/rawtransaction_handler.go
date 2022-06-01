package handlers

import (
	"chainer/internal/app/chainer_app/consumers"
	"chainer/internal/app/chainer_app/models"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RawTransaction struct {
	consumers.RawTransactionConsumer
	Msg   chan amqp.Delivery
	ChErr chan error
}

func NewRawTransaction(consumer consumers.RawTransactionConsumer, msg chan amqp.Delivery, chErr chan error) *RawTransaction {
	return &RawTransaction{
		consumer,
		msg,
		chErr,
	}
}

func (r *RawTransaction) Parser() {

	for {
		select {
		case msg := <-r.Msg:
			r.notify(r.txParser(msg.Body))
			msg.Ack(true)
		case <-r.ChErr:
			fmt.Println("erro returnd")
			//ctx.done
		}
	}
}

func Parse(tx models.Tx) {

	if tx.isWithdraw {
		NotifyDeposit(models.ParsedTx{})

	} else {
		NotifyWithdraw(models.ParsedTx{})
	}

}

func (r *RawTransaction) txParser(msg []byte) *models.Tx {
	tx := &models.Tx{}
	err := json.Unmarshal(msg, tx)
	if err != nil {
		fmt.Printf("faild to marshal message with this err: %s", err)
	}
	return tx
}

func (r *RawTransaction) notify(p *models.Tx) {
	if true {
		fmt.Println("send as deposite transaction")
	}
	fmt.Println("send as withdraw transaction")
}
