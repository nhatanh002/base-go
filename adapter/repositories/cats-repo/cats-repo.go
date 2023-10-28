package cats_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type catsRepo struct {
	db *gorm.DB
}

func NewCatsRepo(db *gorm.DB) implFor.CatsRepository {
	return &catsRepo{db: db}
}

func (r *catsRepo) StoreCat(ctx context.Context, cat model.Cat) error {
	return r.db.Create(&cat).Error
}

func (r *catsRepo) RetrieveCat(ctx context.Context, id string) (*model.Cat, error) {
	cat := model.Cat{}
	err := r.db.Where("id = ?", id).First(&cat).Error
	return &cat, err
}
