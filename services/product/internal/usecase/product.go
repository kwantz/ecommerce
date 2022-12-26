package usecase

import (
	"context"
	"log"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductUsecase struct {
	productRepository ProductRepository
}

func NewProductUsecase(productRepository ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepository: productRepository,
	}
}

func (usecase *ProductUsecase) CreateProduct(ctx context.Context, request entity.ProductRequest) (*entity.ProductResponse, error) {
	operation := "ProductUsecase.CreateProduct"

	product, err := usecase.productRepository.InsertProduct(ctx, request)
	if err != nil {
		log.Printf("[%s] failed insert product from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}, nil
}
