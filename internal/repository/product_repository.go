package repository

import (
	"context"

	"github.com/adi117/Golang-Exercise/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger, db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB:  db,
		Log: log,
	}
}

func (p *ProductRepository) Save(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := p.DB.Create(product).Error
	if err != nil {
		p.Log.Errorf("❌ Failed to save product: %v", err)
		return nil, err
	}
	return product, nil
}

func (p *ProductRepository) GetAll(db *gorm.DB, ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := p.DB.WithContext(ctx).Find(&products).Error
	if err != nil {
		p.Log.Error("❌ Failed to get all products", err)
		return nil, err
	}
	return products, nil
}
