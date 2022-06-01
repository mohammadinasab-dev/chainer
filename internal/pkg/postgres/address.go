package postgres

import (
	"chainer/internal/app/chainer_app/models"
	repo "chainer/internal/app/chainer_app/repositories"
	"context"

	"github.com/go-pg/pg/v10"
)

type Address struct {
	db *pg.DB
}

func NewAddressRepository(db *pg.DB) repo.AddressRepository {
	return &Address{
		db: db,
	}
}

func (a *Address) Insert(ctx context.Context, address *models.Address) error {
	_, err := a.db.WithContext(ctx).Model(address).Insert()
	return err
}
