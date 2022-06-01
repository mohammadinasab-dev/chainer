package models

import "chainer/internal/app/chainer_app/dto"

type Address struct {
	ID       int64 `pg:"id,pk"`
	CoinType string
	AddrStr  string `pg:",unique,notnull"`
}

func NewAddress(id uint32, coinType, address string) *Address {
	return &Address{
		ID:       int64(id),
		CoinType: coinType,
		AddrStr:  address,
	}
}

func (a *Address) ToEntity() *dto.Address {
	return &dto.Address{
		CoinType: a.CoinType,
		AddrStr:  a.AddrStr,
	}
}
