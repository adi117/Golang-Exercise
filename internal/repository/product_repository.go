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
	err := db.Create(product).Error
	if err != nil {
		p.Log.Errorf("❌ Failed to save product: %v", err)
		return nil, err
	}
	return product, nil
}

func (p *ProductRepository) GetAll(db *gorm.DB, ctx context.Context, limit, offset int) ([]*entity.Product, int64, error) {
	var products []*entity.Product
	var total int64

	if err := db.WithContext(ctx).Model(&entity.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.WithContext(ctx).Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		p.Log.Error("❌ Failed to get all products", err)
		return nil, 0, err
	}

	return products, total, nil
}

func (p *ProductRepository) GetByID(db *gorm.DB, id int64) (*entity.Product, error) {
	var product entity.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
