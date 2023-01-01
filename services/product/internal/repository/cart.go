package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (repository *CartRepository) InsertCart(ctx context.Context, request entity.CartRequest) (*entity.Cart, error) {
	operation := "CartRepository.InsertCart"
	query := `
		INSERT cart (account_id, product_id, quantity)
		VALUES (?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		request.AccountID,
		request.ProductID,
		request.Quantity,
	)
	if err != nil {
		log.Printf("[%s] failed execute insert cart, cause: %s", operation, err.Error())
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[%s] failed get last insert ID, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.Cart{
		ID:        insertID,
		AccountID: request.AccountID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}, nil
}

func (repository *CartRepository) DeleteCart(ctx context.Context, cart entity.Cart) (*entity.Cart, error) {
	operation := "CartRepository.DeleteCart"
	query := `
		UPDATE cart 
		SET deleted_at = ?
		WHERE id = ?
	`

	_, err := repository.db.ExecContext(
		ctx,
		query,
		cart.DeletedAt,
		cart.ID,
	)
	if err != nil {
		log.Printf("[%s] failed execute delete cart, cause: %s", operation, err.Error())
		return nil, err
	}

	return &cart, nil
}

func (repository *CartRepository) GetCartByID(ctx context.Context, id int64) (*entity.Cart, error) {
	operation := "CartRepository.GetCartByID"
	query := `
		SELECT id, account_id, product_id, quantity
		FROM cart WHERE id = ? AND deleted_at IS NULL
	`

	result := repository.db.QueryRowContext(ctx, query, id)
	cart := entity.Cart{}
	err := result.Scan(
		&cart.ID,
		&cart.AccountID,
		&cart.ProductID,
		&cart.Quantity,
	)
	if err != nil {
		log.Printf("[%s] failed scan cart result, cause: %s", operation, err.Error())
		return nil, err
	}

	return &cart, nil
}

func (repository *CartRepository) GetCartByAccountID(ctx context.Context, accountID int64) ([]entity.Cart, error) {
	operation := "CartRepository.GetCartByAccountID"
	query := `
		SELECT id, account_id, product_id, quantity
		FROM cart WHERE account_id = ? AND deleted_at IS NULL
	`

	results, err := repository.db.QueryContext(ctx, query, accountID)
	if err != nil {
		log.Printf("[%s] failed query cart, cause: %s", operation, err.Error())
		return nil, err
	}

	carts := []entity.Cart{}

	for results.Next() {
		cart := entity.Cart{}
		err := results.Scan(
			&cart.ID,
			&cart.AccountID,
			&cart.ProductID,
			&cart.Quantity,
		)
		if err != nil {
			log.Printf("[%s] failed scan cart result, cause: %s", operation, err.Error())
			return nil, err
		}

		carts = append(carts, cart)
	}

	return carts, nil
}

func (repository *CartRepository) DeleteCartByAccountID(ctx context.Context, cart entity.Cart) error {
	operation := "CartRepository.DeleteCartByAccountID"
	query := `
		UPDATE cart 
		SET deleted_at = ?
		WHERE account_id = ?
	`

	_, err := repository.db.ExecContext(
		ctx,
		query,
		cart.DeletedAt,
		cart.AccountID,
	)
	if err != nil {
		log.Printf("[%s] failed execute delete cart, cause: %s", operation, err.Error())
		return err
	}

	return nil
}
