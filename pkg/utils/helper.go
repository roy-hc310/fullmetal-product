package utils

import (
	"context"

	"github.com/google/uuid"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
)

func GetUserID(ctx context.Context) *uuid.UUID {
	if v := ctx.Value(constant.UserID); v != nil {
		return v.(*uuid.UUID)
	}
	return nil
}
