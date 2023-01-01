package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderUsecase interface {
	OrderProduct(context.Context, string, int64) (*entity.Order, error)
}

type AccountUsecase interface {
	Authorization(context.Context, string) (int64, error)
}
