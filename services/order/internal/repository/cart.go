package repository

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type CartRepository struct {
	cartHostname string
}

func NewCartRepository(cartHostname string) *CartRepository {
	return &CartRepository{
		cartHostname: cartHostname,
	}
}

func (repository *CartRepository) GetCartProducts(ctx context.Context, auth string) ([]entity.Cart, error) {
	operation := "CartRepository.GetCart"

	req, err := http.NewRequest("POST", repository.cartHostname+"/cart/order", nil)
	if err != nil {
		log.Printf("[%s] failed create new request to get cart order API, cause: %s", operation, err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s] failed do request to get cart order API, cause: %s", operation, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	cart := []entity.Cart{}
	if err := json.NewDecoder(resp.Body).Decode(&cart); err != nil {
		log.Printf("[%s] failed decode cart order API response to get cart, cause: %s", operation, err.Error())
		return nil, err
	}

	return cart, nil
}

func (repository *CartRepository) DeleteCartProducts(ctx context.Context, auth string) ([]entity.Cart, error) {
	operation := "CartRepository.GetCart"

	req, err := http.NewRequest("DELETE", repository.cartHostname+"/cart/order", nil)
	if err != nil {
		log.Printf("[%s] failed create new request to delete cart order API, cause: %s", operation, err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s] failed do request to delete cart order API, cause: %s", operation, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	cart := []entity.Cart{}
	if err := json.NewDecoder(resp.Body).Decode(&cart); err != nil {
		log.Printf("[%s] failed decode cart order API response to delete cart, cause: %s", operation, err.Error())
		return nil, err
	}

	return cart, nil
}
