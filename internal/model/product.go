package model

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" validate:"required"`
	Stock       int      `json:"stock" validate:"required"`
	Category    string   `json:"category"`
	Discount    *float64 `json:"discount"`
}

type CreateProductResponse struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Category    string   `json:"category"`
	Discount    *float64 `json:"discount"`
	CreatedAt   string   `json:"created_at"`
}
