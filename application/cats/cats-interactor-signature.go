package cats

import (
	"base-go/domain/model"
	"context"
)

type CatsService interface {
	AddCat(ctx context.Context, cat model.Cat) error
	GetCat(ctx context.Context, id string) (*model.Cat, error)
}

type AddCatIpt struct {
	Name string `json:"name"`
}

type GetCatResp = model.Cat
