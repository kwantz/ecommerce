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
		INSERT orders (account_id, invoice, status, payment_status, shipping_status)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		order.AccountID,
		order.Invoice,
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

func (repository *OrderRepository) GetAllOrderByAccountID(ctx context.Context, accountID int64) ([]entity.Order, error) {
	operation := "OrderRepository.GetAllOrderByAccountID"
	query := `
		SELECT id, account_id, invoice, status, payment_status, shipping_status
		FROM orders WHERE account_id = ?
	`

	results, err := repository.db.QueryContext(ctx, query, accountID)
	if err != nil {
		log.Printf("[%s] failed query order, cause: %s", operation, err.Error())
		return nil, err
	}

	orders := []entity.Order{}

	for results.Next() {
		order := entity.Order{}
		err := results.Scan(
			&order.ID,
			&order.AccountID,
			&order.Invoice,
			&order.Status,
			&order.PaymentStatus,
			&order.ShippingStatus,
		)
		if err != nil {
			log.Printf("[%s] failed scan order result, cause: %s", operation, err.Error())
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (repository *OrderRepository) GetOrderByAccountIDAndInvoice(ctx context.Context, accountID int64, invoice string) (*entity.Order, error) {
	operation := "OrderRepository.GetOrderByAccountIDAndInvoice"
	query := `
		SELECT id, account_id, invoice, status, payment_status, shipping_status
		FROM orders WHERE account_id = ? AND invoice = ?
	`

	result := repository.db.QueryRowContext(ctx, query, accountID, invoice)
	order := entity.Order{}
	err := result.Scan(
		&order.ID,
		&order.AccountID,
		&order.Invoice,
		&order.Status,
		&order.PaymentStatus,
		&order.ShippingStatus,
	)
	if err != nil {
		log.Printf("[%s] failed scan order result, cause: %s", operation, err.Error())
		return nil, err
	}

	return &order, nil
}
