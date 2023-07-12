package requires

import (
	"base-go/domain/model"
	"context"
)

// dependency signature
type CatsRepository interface {
	StoreCat(ctx context.Context, cat model.Cat) error
	RetrieveCat(ctx context.Context, id string) (*model.Cat, error)
}
