package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderUsecase interface {
	OrderProduct(context.Context, string, int64) (*entity.OrderResponse, error)
	GetOrders(context.Context, int64) ([]entity.OrderResponse, error)
	GetOrderDetail(context.Context, int64, string) (*entity.OrderResponse, error)
}

type AccountUsecase interface {
	Authorization(context.Context, string) (int64, error)
}
