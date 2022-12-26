package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductUsecase interface {
	CreateProduct(context.Context, entity.ProductRequest) (*entity.ProductResponse, error)
	GetAllProduct(context.Context) ([]entity.ProductResponse, error)
}
