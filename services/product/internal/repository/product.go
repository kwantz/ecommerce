package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kwantz/ecommerce/services/product/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repository *ProductRepository) InsertProduct(ctx context.Context, request entity.ProductRequest) (*entity.Product, error) {
	operation := "ProductRepository.InsertProduct"
	query := `
		INSERT product (name, stock, price)
		VALUES (?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		request.Name,
		request.Stock,
		request.Price,
	)
	if err != nil {
		log.Printf("[%s] failed execute insert product, cause: %s", operation, err.Error())
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[%s] failed get last insert ID, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.Product{
		ID:    insertID,
		Name:  request.Name,
		Stock: request.Stock,
		Price: request.Price,
	}, nil
}
