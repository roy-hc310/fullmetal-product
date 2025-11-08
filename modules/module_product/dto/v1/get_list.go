package dto_v1

import (
	"github.com/roy-hc310/fullmetal-product/pkg/base"
)

type GetListProductRequest struct {
	base.Query
	ShopID *string `form:"shop_id" json:"shop_id"`
}

type GetListProductResponse struct {
	base.Entity
	ProductID   string `json:"product_id"`
	BrandID     string `json:"brand_id"`
	CategoryID  string `json:"category_id"`
	ShopID      string `json:"shop_id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	Status      int64  `json:"status"`
}

type PaginationResponse struct {
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
	Total      int64  `json:"total"`
}
