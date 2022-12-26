package usecase

import (
	"context"
	"errors"
	"log"
)

type AccountUsecase struct {
	accountRepository AccountRepository
}

func NewAccountUsecase(accountRepository AccountRepository) *AccountUsecase {
	return &AccountUsecase{
		accountRepository: accountRepository,
	}
}

func (usecase *AccountUsecase) Authorization(ctx context.Context, auth string) (int64, error) {
	operation := "AccountUsecase.Authorization"

	account, err := usecase.accountRepository.GetAccount(ctx, auth)
	if err != nil {
		log.Printf("[%s] failed get account, cause: %s", operation, err.Error())
		return -1, err
	}

	if account.ID <= 0 {
		log.Printf("[%s] invalid account", operation)
		return -1, errors.New("bad request: invalid account")
	}

	return account.ID, nil
}
