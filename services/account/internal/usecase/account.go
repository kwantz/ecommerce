package usecase

import (
	"context"
	"log"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type AccountUsecase struct {
	accountRepository AccountRepository
}

func NewAccountUsecase(accountRepository AccountRepository) *AccountUsecase {
	return &AccountUsecase{
		accountRepository: accountRepository,
	}
}

func (usecase *AccountUsecase) CreateAccount(ctx context.Context, request entity.AccountRequest) (*entity.AccountResponse, error) {
	operation := "AccountUsecase.CreateAccount"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[%s] failed hashed password using bcrypt, cause: %s", operation, err.Error())
		return nil, err
	}

	request.Password = string(hashedPassword)
	account, err := usecase.accountRepository.InsertAccount(ctx, request)
	if err != nil {
		log.Printf("[%s] failed insert account from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	response := convertAccountToAccountResponse(account)
	return &response, nil
}

func (usecase *AccountUsecase) GetAllAccount(ctx context.Context) ([]entity.AccountResponse, error) {
	operation := "AccountUsecase.GetAllAccount"

	accounts, err := usecase.accountRepository.GetAllAccount(ctx)
	if err != nil {
		log.Printf("[%s] failed get all account from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	response := []entity.AccountResponse{}
	for _, account := range accounts {
		response = append(response, convertAccountToAccountResponse(&account))
	}

	return response, nil
}
