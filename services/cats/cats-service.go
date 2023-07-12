package cats

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type catsServiceImpl struct {
	catRepo requires.CatsRepository
}

func NewCatsService(catRepo requires.CatsRepository) implFor.CatsService {
	return &catsServiceImpl{catRepo}
}

func (svc *catsServiceImpl) AddCat(ctx context.Context, cat model.Cat) error {
	return svc.catRepo.StoreCat(ctx, cat)
}

func (svc *catsServiceImpl) GetCat(ctx context.Context, id string) (*model.Cat, error) {
	return svc.catRepo.RetrieveCat(ctx, id)
}
