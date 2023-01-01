package entity

type Cart struct {
	ID           int64 `json:"id"`
	ProductID    int64 `json:"product_id"`
	ProductPrice int64 `json:"product_price"`
	Quantity     int64 `json:"quantity"`
}
