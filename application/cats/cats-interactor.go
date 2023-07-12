package cats

import (
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"

	"github.com/google/uuid"
)

type CatsInteractor struct {
	catService CatsService
}

func NewCatsInteractor(catService CatsService) CatsInteractor {
	return CatsInteractor{catService}
}

func (interactor *CatsInteractor) AddCat(ctx context.Context, cat AddCatIpt) (*model.Cat, error) {
	now := time.Now()
	newCat := model.Cat{
		Id:        uuid.NewString(),
		Name:      cat.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := interactor.catService.AddCat(ctx, newCat)
	if err != nil {
		logger.Error("Unable to create new cat, error: %s", err.Error())
		return nil, err
	}
	return &newCat, nil
}

func (interactor *CatsInteractor) GetCat(ctx context.Context, id string) (*GetCatResp, error) {
	return interactor.catService.GetCat(ctx, id)
}
