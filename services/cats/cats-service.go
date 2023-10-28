package cats

import (
	cats_repo "base-go/adapter/repositories/cats-repo"
	"base-go/domain/model"
	"context"
)

type CatsServiceImpl struct {
	catRepo *cats_repo.CatsRepo
}

func NewCatsService(catRepo *cats_repo.CatsRepo) *CatsServiceImpl {
	return &CatsServiceImpl{catRepo}
}

func (svc *CatsServiceImpl) AddCat(ctx context.Context, cat model.Cat) error {
	return svc.catRepo.StoreCat(ctx, cat)
}

func (svc *CatsServiceImpl) GetCat(ctx context.Context, id string) (*model.Cat, error) {
	return svc.catRepo.RetrieveCat(ctx, id)
}
