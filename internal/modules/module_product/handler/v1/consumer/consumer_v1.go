package consumer_v1

import (
	"context"
	"encoding/json"
	"fmt"

	dto_v1 "github.com/roy-hc310/fullmetal-product/internal/modules/module_product/dto/v1"
	"github.com/roy-hc310/fullmetal-product/internal/modules/module_product/service"
	"github.com/roy-hc310/fullmetal-product/pkg/constant"
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
		data := &dto_v1.CreateProductRequest{}

		if err := json.Unmarshal(value, data); err != nil {
			return err
		}
		_, _, err := h.ProductService.CreateProduct(ctx, data)
		if err != nil {
			return err
		}

		return nil
	// Add more topic handlers here as needed
	// case constant.ProductCreatedTopic:
	//     return h.handleProductCreated(ctx, value)
	case constant.ProductCreateTopic:
		data := &dto_v1.CreateProductRequest{}

		if err := json.Unmarshal(value, data); err != nil {
			return err
		}

		fmt.Println("consume duoc roi ne")
		fmt.Println("Create Product:", data)

		// _, _, err := h.ProductService.CreateProduct(ctx, data)
		// if err != nil {
		// 	return err
		// }
		return nil
	default:
		return fmt.Errorf("unknown topic: %s", topic)
	}
}
