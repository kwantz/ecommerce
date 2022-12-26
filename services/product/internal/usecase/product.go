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

func (usecase *ProductUsecase) GetAllProduct(ctx context.Context) ([]entity.ProductResponse, error) {
	operation := "ProductUsecase.GetAllProduct"

	products, err := usecase.productRepository.GetAllProduct(ctx)
	if err != nil {
		log.Printf("[%s] failed get all product from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	response := []entity.ProductResponse{}
	for _, product := range products {
		response = append(response, entity.ProductResponse{
			ID:    product.ID,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		})
	}

	return response, nil
}

func (usecase *ProductUsecase) GetProduct(ctx context.Context, productID int64) (*entity.ProductResponse, error) {
	operation := "ProductUsecase.GetAllProduct"

	product, err := usecase.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		log.Printf("[%s] failed get product by ID from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}, nil
}
