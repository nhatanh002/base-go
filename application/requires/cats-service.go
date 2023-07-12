package requires

import (
	"base-go/domain/model"
	"context"
)

type CatsService interface {
	AddCat(ctx context.Context, cat model.Cat) error
	GetCat(ctx context.Context, id string) (*model.Cat, error)
}
