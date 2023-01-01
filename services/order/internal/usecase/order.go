package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kwantz/ecommerce/services/order/internal/constant"
	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderUsecaseOption struct {
	CartRepository         CartRepository
	OrderRepository        OrderRepository
	ProductRepository      ProductRepository
	OrderProductRepository OrderProductRepository
}

type OrderUsecase struct {
	cartrepository         CartRepository
	orderRepository        OrderRepository
	productRepository      ProductRepository
	orderProductRepository OrderProductRepository
}

func NewOrderUsecase(option OrderUsecaseOption) *OrderUsecase {
	return &OrderUsecase{
		cartrepository:         option.CartRepository,
		orderRepository:        option.OrderRepository,
		productRepository:      option.ProductRepository,
		orderProductRepository: option.OrderProductRepository,
	}
}

func (usecase *OrderUsecase) OrderProduct(ctx context.Context, auth string, accountID int64) (*entity.OrderResponse, error) {
	operation := "OrderUsecase.OrderProduct"

	cartProducts, err := usecase.cartrepository.GetCartProducts(ctx, auth)
	if err != nil {
		log.Printf("[%s] failed get cart products, cause: %s", operation, err.Error())
		return nil, err
	}

	order, err := usecase.orderRepository.InsertOrder(ctx, entity.Order{
		AccountID:      accountID,
		Invoice:        uuid.New().String(),
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

	return &entity.OrderResponse{
		Invoice:        order.Invoice,
		Status:         order.Status,
		PaymentStatus:  order.PaymentStatus,
		ShippingStatus: order.ShippingStatus,
	}, nil
}

func (usecase *OrderUsecase) GetOrders(ctx context.Context, accountID int64) ([]entity.OrderResponse, error) {
	operation := "OrderUsecase.GetOrders"

	orders, err := usecase.orderRepository.GetAllOrderByAccountID(ctx, accountID)
	if err != nil {
		log.Printf("[%s] failed get all order by account ID, cause: %s", operation, err.Error())
		return nil, err
	}

	response := []entity.OrderResponse{}
	for _, order := range orders {
		response = append(response, entity.OrderResponse{
			Invoice:        order.Invoice,
			Status:         order.Status,
			PaymentStatus:  order.PaymentStatus,
			ShippingStatus: order.ShippingStatus,
		})
	}

	return response, nil
}

func (usecase *OrderUsecase) GetOrderDetail(ctx context.Context, accountID int64, invoice string) (*entity.OrderResponse, error) {
	operation := "OrderUsecase.GetOrderDetail"

	order, err := usecase.orderRepository.GetOrderByAccountIDAndInvoice(ctx, accountID, invoice)
	if err != nil {
		log.Printf("[%s] failed get order by account ID and invoice, cause: %s", operation, err.Error())
		return nil, err
	}

	orderResponse := entity.OrderResponse{
		Invoice:        order.Invoice,
		Status:         order.Status,
		PaymentStatus:  order.PaymentStatus,
		ShippingStatus: order.ShippingStatus,
	}

	orderProducts, err := usecase.orderProductRepository.GetAllOrderProductByOrderID(ctx, order.ID)
	if err != nil {
		log.Printf("[%s] failed get all order product by order ID, cause: %s", operation, err.Error())
		return nil, err
	}

	for _, orderProduct := range orderProducts {
		product, err := usecase.productRepository.GetProductByID(ctx, orderProduct.ProductID)
		if err != nil {
			log.Printf("[%s] failed get get product, cause: %s", operation, err.Error())
			return nil, err
		}
		orderResponse.OrderProductResponse = append(orderResponse.OrderProductResponse, entity.OrderProductResponse{
			ID:          orderProduct.ID,
			ProductID:   orderProduct.ProductID,
			ProductName: product.Name,
			Quantity:    orderProduct.Quantity,
			Price:       orderProduct.Price,
		})
	}

	return &orderResponse, nil
}
