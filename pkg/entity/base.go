package entity

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
