package entity

type OrderProduct struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	Price     int64
}

type OrderProductResponse struct {
	ID          int64  `json:"id"`
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
}
