package cats_repo

import (
	"base-go/domain/model"
	"base-go/services"
	"context"

	"gorm.io/gorm"
)

type catsRepo struct {
	db *gorm.DB
}

func NewCatsRepo(db *gorm.DB) services.CatsRepository {
	return &catsRepo{db: db}
}

func (r *catsRepo) StoreCat(ctx context.Context, cat model.Cat) error {
	return r.db.Create(&cat).Error
}

func (r *catsRepo) RetrieveCat(ctx context.Context, id string) (*model.Cat, error) {
	cat := model.Cat{}
	err := r.db.First(&cat).Where("id = ?", id).Error
	return &cat, err
}
