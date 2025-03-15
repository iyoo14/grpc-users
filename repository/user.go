package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var orderTypeMap = map[int]string{
	1: "ASC",
	2: "DESC",
}

var orderMap = map[int]string{
	1: "id",
	2: "name",
	3: "email",
	4: "age",
	5: "created_at",
}

type UserRepository interface {
	ListUser(ctx context.Context, id int, order int, limit int, orderType int) (UserListEntity, error)
}
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) ListUser(ctx context.Context, id int, order int, limit int, orderType int) (UserListEntity, error) {
	var record []interface{}

	var i = 1
	query := "SELECT id, name, email, age FROM users "
	if id != 0 {
		query += fmt.Sprintf(" WHERE id = $%d", i)
		i++
		record = append(record, id)
	}
	query = query + fmt.Sprintf(" order by %s %s limit $%d", orderMap[order], orderTypeMap[orderType], i)
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
