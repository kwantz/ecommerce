package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
)

type AccountUsecase interface {
	CreateAccount(context.Context, entity.AccountRequest) (*entity.AccountResponse, error)
	GetAllAccount(context.Context) ([]entity.AccountResponse, error)
}

type AuthUsecase interface {
	Authentication(context.Context, entity.LoginAccountRequest) (*entity.LoginAccountResponse, error)
	Authorization(ctx context.Context, tokenString string) (*entity.AccountResponse, error)
}
