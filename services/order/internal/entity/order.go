package entity

type Order struct {
	ID             int64
	AccountID      int64
	Status         string
	PaymentStatus  string
	ShippingStatus string
}

type OrderResponse struct {
	ID                   int64                `json:"id"`
	Status               string               `json:"status"`
	PaymentStatus        string               `json:"payment_status"`
	ShippingStatus       string               `json:"shipping_status"`
	OrderProductResponse OrderProductResponse `json:"list_order_product,omitempty"`
}
