package product_repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/entity"
	"gorm.io/gorm"
)

func (r *ProductRepository) GetDetailProduct(ctx context.Context, id string) (res *dto_v1.GetDetailProductResponse, err error) {

	parseID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	product := entity.Product{}

	err = r.PostgresInfra.DB.WithContext(ctx).Where("id = ?", parseID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, err
	}

	res = &dto_v1.GetDetailProductResponse{}
	if err := copier.Copy(res, &product); err != nil {
		return nil, err
	}

	return res, nil
}

// ✅ Optional: Preload Related Data

// If your Product entity has relationships:

// Brand   Brand   `gorm:"foreignKey:BrandID"`
// Shop    Shop    `gorm:"foreignKey:ShopID"`

// You can preload them like this:

// r.Read.WithContext(ctx).
//     Preload("Brand").
//     Preload("Shop").
//     First(&product, "id = ?", parseID)

// Would you like me to show how to migrate your ListProducts (with filters and pagination) next?
// That’s where GORM query chaining (Where, Limit, Offset, Order) really shines.
