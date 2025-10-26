package dto_v1

import (
	"time"
)

type GetDetailProductResponse struct {
	ProductID   string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	BrandID     string
	CategoryID  string
	ShopID      string
	Name        string
	Sku         string
	Description string
	IsActive    bool
	Status      int64
}
