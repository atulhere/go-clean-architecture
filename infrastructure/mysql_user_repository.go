package infrastructure

import (
	"database/sql"
	"errors"
	"go-clean-architecture/entity"
	"go-clean-architecture/repository"
)

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) repository.UserRepository {
	return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password FROM user WHERE username = ?"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
