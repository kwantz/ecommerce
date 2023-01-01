package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repository *OrderRepository) InsertOrder(ctx context.Context, order entity.Order) (*entity.Order, error) {
	operation := "OrderRepository.InsertOrder"
	query := `
		INSERT orders (account_id, status, payment_status, shipping_status)
		VALUES (?, ?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		order.AccountID,
		order.Status,
		order.PaymentStatus,
		order.ShippingStatus,
	)
	if err != nil {
		log.Printf("[%s] failed execute insert order, cause: %s", operation, err.Error())
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[%s] failed get last insert ID, cause: %s", operation, err.Error())
		return nil, err
	}

	order.ID = insertID
	return &order, nil
}
