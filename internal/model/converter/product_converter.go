package converter

import (
	"time"

	"github.com/adi117/Golang-Exercise/internal/entity"
	"github.com/adi117/Golang-Exercise/internal/model"
)

func ToProductEntity(req model.CreateProductRequest) entity.Product {
	return entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Discount:    req.Discount,
		CreatedAt:   time.Now(),
	}
}

func ToCreateProductResponse(product entity.Product) model.CreateProductResponse {
	return model.CreateProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Discount:    product.Discount,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
	}
}
