package usecase

import (
	"context"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderRepository interface {
	InsertOrder(context.Context, entity.Order) (*entity.Order, error)
	GetAllOrderByAccountID(context.Context, int64) ([]entity.Order, error)
	GetOrderByAccountIDAndInvoice(context.Context, int64, string) (*entity.Order, error)
}

type OrderProductRepository interface {
	InsertOrderProduct(context.Context, entity.OrderProduct) (*entity.OrderProduct, error)
	GetAllOrderProductByOrderID(context.Context, int64) ([]entity.OrderProduct, error)
}

type CartRepository interface {
	GetCartProducts(context.Context, string) ([]entity.Cart, error)
	DeleteCartProducts(context.Context, string) ([]entity.Cart, error)
}

type ProductRepository interface {
	GetProductByID(context.Context, int64) (*entity.Product, error)
}

type AccountRepository interface {
	GetAccount(context.Context, string) (*entity.Account, error)
}
