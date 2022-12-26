package usecase

import (
	"context"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductRepository interface {
	InsertProduct(context.Context, entity.ProductRequest) (*entity.Product, error)
}
