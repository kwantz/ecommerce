package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/kwantz/ecommerce/services/account/internal/entity"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repository *AccountRepository) InsertAccount(ctx context.Context, request entity.AccountRequest) (*entity.Account, error) {
	operation := "AccountRepository.InsertAccount"
	query := `
		INSERT account (password, email, phone, address)
		VALUES (?, ?, ?, ?)
	`

	result, err := repository.db.ExecContext(
		ctx,
		query,
		request.Password,
		request.Email,
		request.Phone,
		request.Address,
	)
	if err != nil {
		log.Printf("[%s] failed execute insert account, cause: %s", operation, err.Error())
		return nil, err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Printf("[%s] failed get last insert ID, cause: %s", operation, err.Error())
		return nil, err
	}

	return &entity.Account{
		ID:       insertID,
		Password: request.Password,
		Email:    request.Email,
		Phone:    request.Phone,
		Address:  request.Address,
	}, nil
}

func (repository *AccountRepository) GetAllAccount(ctx context.Context) ([]entity.Account, error) {
	operation := "AccountRepository.GetAllAccount"
	query := `
		SELECT id, password, email, phone, address 
		FROM account
	`

	results, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[%s] failed query account, cause: %s", operation, err.Error())
		return nil, err
	}

	accounts := []entity.Account{}

	for results.Next() {
		account := entity.Account{}
		err := results.Scan(
			&account.ID,
			&account.Password,
			&account.Email,
			&account.Phone,
			&account.Address,
		)
		if err != nil {
			log.Printf("[%s] failed scan account result, cause: %s", operation, err.Error())
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (repository *AccountRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	operation := "AccountRepository.GetAccountByEmail"
	query := `
		SELECT id, password, email, phone, address 
		FROM account
		WHERE email = ?
	`

	result := repository.db.QueryRowContext(ctx, query, email)
	account := entity.Account{}
	err := result.Scan(
		&account.ID,
		&account.Password,
		&account.Email,
		&account.Phone,
		&account.Address,
	)
	if err != nil {
		log.Printf("[%s] failed scan account result, cause: %s", operation, err.Error())
		return nil, err
	}

	return &account, nil
}
