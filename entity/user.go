package entity

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetUserByUsername(username string) (*entity.User, error)
	CreateUser(user *entity.User) error
}
