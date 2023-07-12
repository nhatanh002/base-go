package services

import (
	"base-go/application/cats"
	"base-go/domain/model"
	"context"
)

// dependency signature
type CatsRepository interface {
	StoreCat(ctx context.Context, cat model.Cat) error
	RetrieveCat(ctx context.Context, id string) (*model.Cat, error)
}

// impl
type catsServiceImpl struct {
	catRepo CatsRepository
}

func NewCatsService(catRepo CatsRepository) cats.CatsService {
	return &catsServiceImpl{catRepo}
}

func (svc *catsServiceImpl) AddCat(ctx context.Context, cat model.Cat) error {
	return svc.catRepo.StoreCat(ctx, cat)
}

func (svc *catsServiceImpl) GetCat(ctx context.Context, id string) (*model.Cat, error) {
	return svc.catRepo.RetrieveCat(ctx, id)
}
