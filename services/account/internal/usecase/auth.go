package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kwantz/ecommerce/services/account/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	accountRepository AccountRepository
}

func NewAuthUsecase(accountRepository AccountRepository) *AuthUsecase {
	return &AuthUsecase{
		accountRepository: accountRepository,
	}
}

func (usecase *AuthUsecase) Authentication(ctx context.Context, request entity.LoginAccountRequest) (*entity.LoginAccountResponse, error) {
	operation := "AuthUsecase.Authentication"

	account, err := usecase.accountRepository.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		log.Printf("[%s] failed get all account from repository, cause: %s", operation, err.Error())
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		log.Printf("[%s] failed compare hash and password, cause: %s", operation, err.Error())
		return nil, err
	}

	claim := entity.LoginAccountClaim{
		AccountResponse: convertAccountToAccountResponse(account),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
			Issuer:    "ACCOUNT_SERVICE_ISSUER",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte("this_is_secret_key"))
	if err != nil {
		log.Printf("[%s] failed signed token JWT, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.LoginAccountResponse{
		Token: signedToken,
	}, nil
}

func (usecase *AuthUsecase) Authorization(ctx context.Context, tokenString string) (*entity.AccountResponse, error) {
	operation := "AuthUsecase.Authorization"

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing token method")
		}

		return []byte("this_is_secret_key"), nil
	})
	if err != nil {
		log.Printf("[%s] failed parse token JTW, cause: %s", operation, err.Error())
		return nil, err
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Printf("[%s] invalid token claim", operation)
		return nil, errors.New("bad request")
	}

	jsonByte, err := json.Marshal(claim)
	if err != nil {
		log.Printf("[%s] failed marshal claim, cause: %s", operation, err.Error())
		return nil, err
	}

	account := entity.AccountResponse{}
	if err := json.Unmarshal(jsonByte, &account); err != nil {
		log.Printf("[%s] failed unmarshal to account, cause: %s", operation, err.Error())
		return nil, err
	}

	return &account, nil
}
