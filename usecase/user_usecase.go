package usecase

import (
	"errors"
	"go-clean-architecture/repository"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (u *UserUsecase) Login(username string, password string) (string, error) {
	user, err := u.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Here you can hash and compare the password (for simplicity, skipped)
	if user.Password != password {
		return "", errors.New("invalid username or password")
	}

	// Here, generate JWT token (omitted for simplicity)
	return "success", nil
}
