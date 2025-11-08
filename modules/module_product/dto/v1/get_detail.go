package dto_v1

import (
	"github.com/roy-hc310/fullmetal-product/pkg/base"
)

type GetDetailProductResponse struct {
	base.Entity
	BrandID     string `json:"brand_id"`
	CategoryID  string `json:"category_id"`
	ShopID      string `json:"shop_id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	Status      int64  `json:"status"`
}
