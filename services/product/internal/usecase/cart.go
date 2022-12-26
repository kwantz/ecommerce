package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type CartUsecase struct {
	cartRepository    CartRepository
	productRepository ProductRepository
}

func NewCartUsecase(cartRepository CartRepository, productRepository ProductRepository) *CartUsecase {
	return &CartUsecase{
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (usecase *CartUsecase) AddProductToCart(ctx context.Context, request entity.CartRequest) (*entity.CartResponse, error) {
	operation := "CartUsecase.AddProductToCart"

	if request.AccountID <= 0 {
		log.Printf("[%s] invalid account ID", operation)
		return nil, errors.New("invalid account ID")
	}

	product, err := usecase.productRepository.GetProductByID(ctx, request.ProductID)
	if err != nil {
		log.Printf("[%s] failed get product by ID, cause: %s", operation, err.Error())
		return nil, err
	}

	if product.Stock < request.Quantity {
		log.Printf("[%s] failed cause quantity greater than stock product", operation)
		return nil, errors.New("quantity greater than stock product")
	}

	product.Stock -= request.Quantity

	if _, err := usecase.productRepository.UpdateProduct(ctx, *product); err != nil {
		log.Printf("[%s] failed update product, cause: %s", operation, err.Error())
		return nil, err
	}

	cart, err := usecase.cartRepository.InsertCart(ctx, request)
	if err != nil {
		log.Printf("[%s] failed insert cart, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.CartResponse{
		ID:        cart.ID,
		AccountID: cart.AccountID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
	}, nil
}

func (usecase *CartUsecase) RemoveProductFromCart(ctx context.Context, cartID int64) (*entity.CartResponse, error) {
	operation := "CartUsecase.RemoveProductFromCart"

	cart, err := usecase.cartRepository.GetCartByID(ctx, cartID)
	if err != nil {
		log.Printf("[%s] failed get cart by ID, cause: %s", operation, err.Error())
		return nil, err
	}

	product, err := usecase.productRepository.GetProductByID(ctx, cart.ProductID)
	if err != nil {
		log.Printf("[%s] failed get product by ID, cause: %s", operation, err.Error())
		return nil, err
	}

	product.Stock += cart.Quantity

	if _, err := usecase.productRepository.UpdateProduct(ctx, *product); err != nil {
		log.Printf("[%s] failed update product, cause: %s", operation, err.Error())
		return nil, err
	}

	cart.DeletedAt = time.Now()

	if _, err := usecase.cartRepository.DeleteCart(ctx, *cart); err != nil {
		log.Printf("[%s] failed delete cart, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.CartResponse{
		ID:        0,
		AccountID: 0,
		ProductID: 0,
		Quantity:  0,
	}, nil
}
