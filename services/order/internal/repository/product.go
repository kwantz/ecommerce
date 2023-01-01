package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type ProductRepository struct {
	productHostname string
}

func NewProductRepository(productHostname string) *ProductRepository {
	return &ProductRepository{
		productHostname: productHostname,
	}
}

func (repository *ProductRepository) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	operation := "ProductRepository.GetProductByID"

	url := fmt.Sprintf("%s/%d", repository.productHostname, productID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[%s] failed create new request to product API, cause: %s", operation, err.Error())
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s] failed do request to product API, cause: %s", operation, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	product := entity.Product{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		log.Printf("[%s] failed decode product API response to product, cause: %s", operation, err.Error())
		return nil, err
	}

	return &product, nil
}
