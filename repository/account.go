package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type AccountRepository interface {
	GetApiKey(ctx context.Context, accountId int, apiKey string) (AccountEntity, error)
}
type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) GetApiKey(ctx context.Context, accountId int, apiKey string) (AccountEntity, error) {
	var record []interface{}

	fmt.Println(accountId, apiKey)
	query := "SELECT api_key FROM account where account_id = $1 and api_key = $2"
	record = append(record, accountId, apiKey)

	fmt.Println(query, record)
	var account Account
	var accountEntity AccountEntity
	err := r.db.QueryRowx(query, record...).StructScan(&account)
	if err != nil {
		fmt.Println(query, record)
		return accountEntity, err
	}

	fmt.Println(reflect.TypeOf(account.ApiKey))
	accountEntity.ApiKey = account.ApiKey.(string)
	return accountEntity, nil
}
