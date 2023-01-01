package usecase

import (
	"context"
	"log"

	"github.com/kwantz/ecommerce/services/order/internal/constant"
	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderUsecaseOption struct {
	CartRepository         CartRepository
	OrderRepository        OrderRepository
	OrderProductRepository OrderProductRepository
}

type OrderUsecase struct {
	cartrepository         CartRepository
	orderRepository        OrderRepository
	orderProductRepository OrderProductRepository
}

func NewOrderUsecase(option OrderUsecaseOption) *OrderUsecase {
	return &OrderUsecase{
		cartrepository:         option.CartRepository,
		orderRepository:        option.OrderRepository,
		orderProductRepository: option.OrderProductRepository,
	}
}

func (usecase *OrderUsecase) OrderProduct(ctx context.Context, auth string, accountID int64) (*entity.Order, error) {
	operation := "OrderUsecase.OrderProduct"

	cartProducts, err := usecase.cartrepository.GetCartProducts(ctx, auth)
	if err != nil {
		log.Printf("[%s] failed get cart products, cause: %s", operation, err.Error())
		return nil, err
	}

	order, err := usecase.orderRepository.InsertOrder(ctx, entity.Order{
		AccountID:      accountID,
		Status:         constant.ORDER_STATUS_PENDING,
		PaymentStatus:  constant.PAYMENT_STATUS_AWAITING,
		ShippingStatus: constant.SHIPPING_STATUS_AWAITING,
	})
	if err != nil {
		log.Printf("[%s] failed insert order, cause: %s", operation, err.Error())
		return nil, err
	}

	for _, cart := range cartProducts {
		_, err := usecase.orderProductRepository.InsertOrderProduct(ctx, entity.OrderProduct{
			OrderID:   order.ID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.ProductPrice,
		})
		if err != nil {
			log.Printf("[%s] failed insert order product, cause: %s", operation, err.Error())
			return nil, err
		}
	}

	if _, err := usecase.cartrepository.DeleteCartProducts(ctx, auth); err != nil {
		log.Printf("[%s] failed delete cart products, cause: %s", operation, err.Error())
		return nil, err
	}

	return order, nil
}

func (usecase *OrderUsecase) GetOrder() {
	// Get Order By Account ID
	// Return
}

func (usecase *OrderUsecase) GetOrderDetail() {
	// Get Order By ID & Account ID
	// Loop & Get Product By ID
	// Return
}
