package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserRepository interface {
	ListUser(ctx context.Context, order string, limit int) (UserListEntity, error)
}
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) ListUser(ctx context.Context, order string, limit int) (UserListEntity, error) {
	var record []interface{}
	query := fmt.Sprintf("SELECT id, name, email, age FROM users order by %s  limit $1", order)
	record = append(record, limit)
	fmt.Println(query, record)
	userList, err := r.fetchUsers(query, record...)
	if err != nil {
		fmt.Println(query, record)
		log.Fatal(err)
	}
	return userList, err
}

func (r *userRepository) fetchUsers(query string, args ...interface{}) (UserListEntity, error) {
	var userListEntity UserListEntity
	stmt, err := r.db.Preparex(query)
	if err != nil {
		return userListEntity, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Queryx(args...)
	if err != nil {
		return userListEntity, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.StructScan(&user); err != nil {
			return userListEntity, fmt.Errorf("failed to scan row: %w", err)
		}
		fmt.Println(user)
		userListEntity.UserList = append(userListEntity.UserList, user)
	}
	return userListEntity, nil
}
