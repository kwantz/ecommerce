package controller

import (
	"context"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
)

type AccountUsecase interface {
	LoginAccount(context.Context, entity.LoginAccountRequest) (*entity.LoginAccountResponse, error)
	Authorization(ctx context.Context, tokenString string) (*entity.AccountResponse, error)

	CreateAccount(context.Context, entity.AccountRequest) (*entity.AccountResponse, error)
	GetAllAccount(context.Context) ([]entity.AccountResponse, error)
}
