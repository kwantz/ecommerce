package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kwantz/ecommerce/services/order/internal/entity"
)

type OrderProductRepository struct {
	db *sql.DB
}

func NewOrderProductRepository(db *sql.DB) *OrderProductRepository {
	return &OrderProductRepository{
		db: db,
	}
}

func (repository *OrderProductRepository) InsertOrderProduct(ctx context.Context, orderProduct entity.OrderProduct) (*entity.OrderProduct, error) {
	operation := "OrderProductRepository.InsertOrderProduct"
	query := `
		INSERT order_product (order_id, product_id, quantity, price)
		VALUES (?, ?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		orderProduct.OrderID,
		orderProduct.ProductID,
		orderProduct.Quantity,
		orderProduct.Price,
	)
	if err != nil {
		log.Printf("[%s] failed execute insert order product, cause: %s", operation, err.Error())
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[%s] failed get last insert ID, cause: %s", operation, err.Error())
		return nil, err
	}

	orderProduct.ID = insertID
	return &orderProduct, nil
}

func (repository *OrderProductRepository) GetAllOrderProductByOrderID(ctx context.Context, orderID int64) ([]entity.OrderProduct, error) {
	operation := "OrderProductRepository.GetAllOrderProductByOrderID"
	query := `
		SELECT id, order_id, product_id, quantity, price
		FROM order_product WHERE order_id = ?
	`

	results, err := repository.db.QueryContext(ctx, query, orderID)
	if err != nil {
		log.Printf("[%s] failed query order product, cause: %s", operation, err.Error())
		return nil, err
	}

	orderProducts := []entity.OrderProduct{}

	for results.Next() {
		orderProduct := entity.OrderProduct{}
		err := results.Scan(
			&orderProduct.ID,
			&orderProduct.OrderID,
			&orderProduct.ProductID,
			&orderProduct.Quantity,
			&orderProduct.Price,
		)
		if err != nil {
			log.Printf("[%s] failed scan order product result, cause: %s", operation, err.Error())
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}
