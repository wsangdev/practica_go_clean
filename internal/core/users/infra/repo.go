package infra

import (
	"one/internal/core/users/domain"

	"gorm.io/gorm"
)

type UserGorm struct {
	db *gorm.DB
}

func NewUserGorm(database *gorm.DB) domain.UserRepository {
	return &UserGorm{db: database}
}

// 
func (r UserGorm) GetAllUsers() ([]domain.UserEntity, error) {
	var models []UserModel

	result := r.db.Raw("SELECT id, nombre, email, edad, FROM users").Scan(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	return ToDomainList(models), nil
}
