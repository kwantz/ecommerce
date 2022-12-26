package entity

type Account struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
