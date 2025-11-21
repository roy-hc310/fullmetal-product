package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	UserID    *uuid.UUID `db:"user_id" json:"user_id"`
	// BrandID     *uuid.UUID `db:"brand_id" json:"brand_id,omitempty"`
	// CategoryID  *uuid.UUID `db:"category_id" json:"category_id,omitempty"`
	// ShopID      *uuid.UUID `db:"shop_id" json:"shop_id,omitempty"`
	Name        string  `db:"name" json:"name"`
	SKU         *string `db:"sku" json:"sku,omitempty"`
	Description *string `db:"description" json:"description,omitempty"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	Status      int     `db:"status" json:"status"`
}

func (p *Product) Validation() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
