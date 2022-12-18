package usecase

import "github.com/kwantz/ecommerce/services/account/internal/entity"

func (usecase *AccountUsecase) convertAccountToAccountResponse(account *entity.Account) entity.AccountResponse {
	return entity.AccountResponse{
		ID:      account.ID,
		Email:   account.Email,
		Phone:   account.Phone,
		Address: account.Address,
	}
}
