package repositories

import (
	"chainer/internal/app/chainer_app/models"
	"context"
)

type AddressRepository interface {
	Insert(ctx context.Context, address *models.Address) error
}
