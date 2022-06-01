package controllers

import (
	"chainer/internal/app/chainer_app/handlers"
	"chainer/pkg/pbs"
	"context"
)

type Address struct {
	handlers.AddressHandler
}

func NewAddress(handler *handlers.Address) *Address {
	return &Address{
		handler,
	}
}

func (a *Address) GetDepositAddress(ctx context.Context, request *pbs.DepositAddressReq) (respone *pbs.DepositAddressRes, err error) {
	address, err := a.GetDepositAddressHandler(ctx, request.Asset, request.Network)

	if err != nil {
		return nil, err
	}
	return &pbs.DepositAddressRes{
		Address: address.AddrStr,
	}, err

}
