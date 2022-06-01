package handlers

import (
	"chainer/internal/app/chainer_app/models"
	"chainer/internal/app/chainer_app/publishers"
	repo "chainer/internal/app/chainer_app/repositories"
	"context"
	"fmt"
)

type WithdrawHandler interface {
	Withdrawal(ctx context.Context, asset, network, address string, amount int64) (txid string, err error)
}

type Withdraw struct {
	repo.AddressRepository
	publishers.AddressPublisher
}

func NewWithdrawHandler(adr repo.AddressRepository, publisher publishers.AddressPublisher) *Withdraw {
	return &Withdraw{
		adr,
		publisher,
	}
}

func (w *Withdraw) withdrawal(ctx context.Context, asset, network, address string, amount int64) (string, error) {
	fmt.Println("withdraw")

	TXID := "unImplemented"

	return TXID, nil
}

func NotifyWithdraw(models.ParsedTx) {

}
