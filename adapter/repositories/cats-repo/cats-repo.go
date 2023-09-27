package cats_repo

import (
	"base-go/domain/model"
	"context"

	"gorm.io/gorm"
)

type CatsRepo struct {
	db *gorm.DB
}

func NewCatsRepo(db *gorm.DB) *CatsRepo {
	return &CatsRepo{db: db}
}

func (r *CatsRepo) StoreCat(ctx context.Context, cat model.Cat) error {
	return r.db.Create(&cat).Error
}

func (r *CatsRepo) RetrieveCat(ctx context.Context, id string) (*model.Cat, error) {
	cat := model.Cat{}
	err := r.db.First(&cat).Where("id = ?", id).Error
	return &cat, err
}
