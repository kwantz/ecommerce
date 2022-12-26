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

func (repository *ProductRepository) UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error) {
	operation := "ProductRepository.UpdateProduct"
	query := `
		UPDATE product 
		SET name = ?, stock = ?, price = ?
		WHERE id = ?
	`

	_, err := repository.db.ExecContext(
		ctx,
		query,
		product.Name,
		product.Stock,
		product.Price,
		product.ID,
	)
	if err != nil {
		log.Printf("[%s] failed execute update product, cause: %s", operation, err.Error())
		return nil, err
	}

	return &product, nil
}

func (repository *ProductRepository) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	operation := "ProductRepository.GetAllProduct"
	query := `
		SELECT id, name, stock, price
		FROM product
	`

	results, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[%s] failed query product, cause: %s", operation, err.Error())
		return nil, err
	}

	products := []entity.Product{}

	for results.Next() {
		product := entity.Product{}
		err := results.Scan(
			&product.ID,
			&product.Name,
			&product.Stock,
			&product.Price,
		)
		if err != nil {
			log.Printf("[%s] failed scan product result, cause: %s", operation, err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductRepository) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	operation := "ProductRepository.GetProductByID"
	query := `
		SELECT id, name, stock, price
		FROM product WHERE id = ?
	`

	result := repository.db.QueryRowContext(ctx, query, id)
	product := entity.Product{}
	err := result.Scan(
		&product.ID,
		&product.Name,
		&product.Stock,
		&product.Price,
	)
	if err != nil {
		log.Printf("[%s] failed scan product result, cause: %s", operation, err.Error())
		return nil, err
	}

	return &product, nil
}
