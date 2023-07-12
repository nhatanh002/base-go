package services

import (
	"base-go/application/cats"
	"base-go/domain/model"
	"context"
)

// dependency signature
type CatsRepository interface {
	AddCat(ctx context.Context, cat model.Cat) error
	GetCat(ctx context.Context, id string) (*model.Cat, error)
}

// impl
type CatsServiceImpl struct {
	catRepo CatsRepository
}

func NewCatsService(catRepo CatsRepository) cats.CatsService {
	return &CatsServiceImpl{catRepo}
}

var _ cats.CatsService = &CatsServiceImpl{}

func (svc *CatsServiceImpl) AddCat(ctx context.Context, cat model.Cat) error {
	return svc.catRepo.AddCat(ctx, cat)
}

func (svc *CatsServiceImpl) GetCat(ctx context.Context, id string) (*model.Cat, error) {
	return svc.catRepo.GetCat(ctx, id)
}
