package usecase

import (
	"context"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
)

type AccountRepository interface {
	InsertAccount(context.Context, entity.AccountRequest) (*entity.Account, error)
	GetAllAccount(context.Context) ([]entity.Account, error)
	GetAccountByEmail(context.Context, string) (*entity.Account, error)
}
