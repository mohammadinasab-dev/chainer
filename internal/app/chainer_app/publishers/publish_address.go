package publishers

import (
	"context"
)

type AddressPublisher interface {
	Publisher(ctx context.Context, msg []byte) error
}
