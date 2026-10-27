package domain

type UserRepository interface {
	GetAllUsers() ([]UserEntity, error)
	// GetUserById(id int64) (*UserEntity, error)
	// CreateUser(user *UserEntity) error
}
