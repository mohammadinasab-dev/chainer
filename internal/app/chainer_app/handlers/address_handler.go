package handlers

import (
	"chainer/internal/app/chainer_app/crypto"
	"chainer/internal/app/chainer_app/dto"
	"chainer/internal/app/chainer_app/models"
	"chainer/internal/app/chainer_app/publishers"
	repo "chainer/internal/app/chainer_app/repositories"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type AddressHandler interface {
	GetDepositAddressHandler(ctx context.Context, asset, network string) (walletAddres *dto.Address, err error)
}

type Address struct {
	repo.AddressRepository
	publishers.AddressPublisher
}

func NewAddress(adr repo.AddressRepository, publisher publishers.AddressPublisher) *Address {
	return &Address{
		adr,
		publisher,
	}
}

func (c *Address) GetDepositAddressHandler(ctx context.Context, asset, network string) (walletAddres *dto.Address, err error) {
	//generate uuid
	id := generateUUID()
	walletAddress, err := crypto.GenerateParkingAddress(network, id)

	address := models.NewAddress(id, asset, walletAddress)
	//insert into db
	err = c.Insert(ctx, address)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bytes := address.ToEntity().ToMessage().ToBytes()

	c.Publisher(ctx, bytes)

	return &dto.Address{
		AddrStr: address.AddrStr,
	}, err

}

func generateUUID() uint32 {
	uuid := uuid.New()
	return uuid.ID()
}
