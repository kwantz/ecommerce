package entity

import "github.com/golang-jwt/jwt/v4"

type Account struct {
	ID       int64
	Password string
	Email    string
	Phone    string
	Address  string
}

type AccountRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type AccountResponse struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type LoginAccountRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginAccountResponse struct {
	Token string `json:"token"`
}

type LoginAccountClaim struct {
	jwt.RegisteredClaims
	AccountResponse
}
