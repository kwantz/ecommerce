package entity

import "time"

type Cart struct {
	ID        int64
	AccountID int64
	ProductID int64
	Quantity  int64
	DeletedAt time.Time
}

type CartRequest struct {
	AccountID int64
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type CartResponse struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}
