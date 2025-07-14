package usecase

import (
	"context"
	"errors"

	"github.com/adi117/Golang-Exercise/internal/entity"
	"github.com/adi117/Golang-Exercise/internal/model"
	"github.com/adi117/Golang-Exercise/internal/model/converter"
	"github.com/adi117/Golang-Exercise/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.CreateProductResponse, error)
	GetAllProducts(ctx context.Context, limit, offset int) ([]*entity.Product, int64, error)
	GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, id int64, req *model.UpdateProductRequest) (*model.CreateProductResponse, error)
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

func (p *productUsecase) GetAllProducts(ctx context.Context, limit, offset int) ([]*entity.Product, int64, error) {
	products, total, err := p.ProductRepository.GetAll(p.DB, ctx, limit, offset)
	if err != nil {
		p.Log.Error("failed to retrieve all products")
		return nil, 0, err
	}
	return products, total, nil
}

func (p *productUsecase) GetProductByID(ctx context.Context, id int64) (*model.CreateProductResponse, error) {
	product, err := p.ProductRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Product not found")
		}
		p.Log.WithError(err).Error("Failed to get product")
		return nil, err
	}
	response := converter.ToCreateProductResponse(*product)
	return &response, nil
}

func (p *productUsecase) UpdateProduct(ctx context.Context, id int64, req *model.UpdateProductRequest) (*model.CreateProductResponse, error) {

	tx := p.DB.Begin()

	currentProduct, err := p.ProductRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Product not found")
		}

		p.Log.WithError(err).Error("Failed to get product")
		return nil, err
	}

	if req.Name != nil {
		currentProduct.Name = *req.Name
	}

	if req.Category != nil {
		currentProduct.Category = *req.Category
	}

	if req.Description != nil {
		currentProduct.Description = *req.Description
	}

	if req.Discount != nil {
		currentProduct.Discount = req.Discount
	}

	if req.Price != nil {
		currentProduct.Price = *req.Discount
	}

	if req.Price != nil {
		currentProduct.Price = *req.Price
	}

	if req.Stock != nil {
		currentProduct.Stock = *req.Stock
	}

	// save updated product to DB
	updatedProduct, err := p.ProductRepository.Update(tx, currentProduct)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to update product")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	response := converter.ToCreateProductResponse(*updatedProduct)
	return &response, nil
}
