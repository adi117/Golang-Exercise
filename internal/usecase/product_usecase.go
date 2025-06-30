package usecase

import (
	"context"

	"github.com/adi117/Golang-Exercise/internal/entity"
	"github.com/adi117/Golang-Exercise/internal/model"
	"github.com/adi117/Golang-Exercise/internal/model/converter"
	"github.com/adi117/Golang-Exercise/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
}

type productUsecase struct {
	ProductRepository *repository.ProductRepository
	Log               *logrus.Logger
	DB                *gorm.DB
}

func NewProductUsecase(productRepository *repository.ProductRepository, log *logrus.Logger, db *gorm.DB) ProductUsecase {
	return &productUsecase{
		ProductRepository: productRepository,
		Log:               log,
		DB:                db,
	}
}

func (p *productUsecase) CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error) {
	tx := p.DB.Begin()
	product := converter.ToProductEntity(*req)
	savedProduct, err := p.ProductRepository.Save(tx, &product)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to create the product")
		return nil, err
	}

	response := converter.ToCreateProductResponse(*savedProduct)

	return &response, err
}

func (p *productUsecase) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	products, err := p.ProductRepository.GetAll(p.DB, ctx)
	if err != nil {
		p.Log.Error("failed to retrieve all products")
		return nil, err
	}
	return products, nil
}
