package usecase

import (
	"context"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductRepository interface {
	InsertProduct(context.Context, entity.ProductRequest) (*entity.Product, error)
	UpdateProduct(context.Context, entity.Product) (*entity.Product, error)
	GetAllProduct(context.Context) ([]entity.Product, error)
	GetProductByID(context.Context, int64) (*entity.Product, error)
}

type CartRepository interface {
	InsertCart(context.Context, entity.CartRequest) (*entity.Cart, error)
	DeleteCart(context.Context, entity.Cart) (*entity.Cart, error)
	GetCartByID(context.Context, int64) (*entity.Cart, error)

	GetCartByAccountID(context.Context, int64) ([]entity.Cart, error)
	DeleteCartByAccountID(context.Context, entity.Cart) error
}

type AccountRepository interface {
	GetAccount(context.Context, string) (*entity.Account, error)
}
