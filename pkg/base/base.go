package base

import (
	"time"

	"github.com/google/uuid"
)

type Query struct {
	Page   int    `form:"page" json:"page"`
	Limit  int    `form:"limit" json:"limit"`
	Cursor string `form:"cursor" json:"cursor"`
	SortBy string `form:"sort" json:"sort_by"`
	Order  string `form:"order" json:"order"`
}

type Entity struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
