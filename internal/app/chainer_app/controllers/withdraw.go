package controllers

import (
	"chainer/internal/app/chainer_app/handlers"
	"chainer/pkg/pbs"
	"context"
	"fmt"
)

type Withdraw struct {
	handlers.WithdrawHandler
}

func NewWithdraw(handler handlers.WithdrawHandler) *Withdraw {
	return &Withdraw{
		handler,
	}
}

func (a *Withdraw) Withdraw(ctx context.Context, request *pbs.WithdrawReq) (respone *pbs.WithdrawRes, err error) {
	fmt.Println("withdraw")
	txid, err := a.Withdrawal(ctx, request.Asset, request.Network, request.Address, request.Amount)
	if err != nil {
		return nil, err
	}
	respone = &pbs.WithdrawRes{
		TXID: txid,
	}
	return
}
