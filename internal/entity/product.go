package entity

import "time"

type Product struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Price       float64   `gorm:"column:price;type:numeric"`
	Stock       int       `gorm:"column:stock"`
	Category    string    `gorm:"column:category"`
	Discount    *float64  `gorm:"column:discount;type:numeric"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestampz;default:now()"`
}

func (p *Product) TableName() string {
	return "product"
}
