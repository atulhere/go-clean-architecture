package repository

import "go-clean-architecture/entity"

type UserRepository interface {
	GetUserByUsername(username string) (*entity.User, error)
	CreateUser(user *entity.User) error
}
