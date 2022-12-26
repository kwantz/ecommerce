package entity

type Product struct {
	ID    int64
	Name  string
	Stock int64
	Price int64
}

type ProductRequest struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

type ProductResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}
