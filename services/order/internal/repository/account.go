package repository

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type AccountRepository struct {
	accountHostname string
}

func NewAccountRepository(accountHostname string) *AccountRepository {
	return &AccountRepository{
		accountHostname: accountHostname,
	}
}

func (repository *AccountRepository) GetAccount(ctx context.Context, auth string) (*entity.Account, error) {
	operation := "AccountRepository.GetAccount"

	req, err := http.NewRequest("POST", repository.accountHostname+"/authorize", nil)
	if err != nil {
		log.Printf("[%s] failed create new request to authorize API, cause: %s", operation, err.Error())
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s] failed do request to authorize API, cause: %s", operation, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	account := entity.Account{}
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		log.Printf("[%s] failed decode authorize API response to account, cause: %s", operation, err.Error())
		return nil, err
	}

	return &account, nil
}
