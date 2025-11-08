package dto_v1

type CreateProductRequest struct {
	BrandID    string
	CategoryID string
	ShopID     string

	Name        string
	Sku         string
	Description string
	IsActive    bool
	Status      int64
}
