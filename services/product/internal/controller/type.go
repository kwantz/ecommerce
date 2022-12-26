package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductUsecase interface {
	CreateProduct(context.Context, entity.ProductRequest) (*entity.ProductResponse, error)
	GetAllProduct(context.Context) ([]entity.ProductResponse, error)
	GetProduct(context.Context, int64) (*entity.ProductResponse, error)
}

type CartUsecase interface {
	AddProductToCart(context.Context, entity.CartRequest) (*entity.CartResponse, error)
	RemoveProductFromCart(context.Context, int64) (*entity.CartResponse, error)
}

type AccountUsecase interface {
	Authorization(context.Context, string) (int64, error)
}
