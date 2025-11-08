package base

import (
	"time"

	"github.com/google/uuid"
)

type Query struct {
	Page   *string `form:"page" json:"page"`
	Size   *string `form:"size" json:"size"`
	Cursor *string `form:"cursor" json:"cursor"`
	SortBy *string `form:"sort" json:"sort_by"`
}

type Entity struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
