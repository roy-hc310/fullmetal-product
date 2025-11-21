package consumer_v1

import (
	"context"
	"fmt"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
	"github.com/roy-hc310/fullmetal-product/pkg/logger"
	"github.com/roy-hc310/fullmetal-product/pkg/utils"
)

type ProductConsumer struct {
	ProductService service.ProductServiceInterface
}

func NewProductConsumer(productService service.ProductServiceInterface) *ProductConsumer {
	return &ProductConsumer{
		ProductService: productService,
	}
}

func (h *ProductConsumer) HandleMessage(ctx context.Context, topic string, value []byte) error {
	switch topic {
	case constant.DefaultTopic:
		logger.Info(ctx).Str("topic", topic).Msg("Received message on default topic")
		return nil
	case constant.ProductCreateTopic:
		data := &dto_v1.CreateProductRequest{}

		if err := utils.JSONToStruct(value, data); err != nil {
			return err
		}

		logger.Info(ctx).
			Str("topic", topic).
			Interface("data", data).
			Msg("Received product create event")

		// _, _, err := h.ProductService.CreateProduct(ctx, data)
		// if err != nil {
		// 	return err
		// }
		return nil
	default:
		return fmt.Errorf("unknown topic: %s", topic)
	}
}
