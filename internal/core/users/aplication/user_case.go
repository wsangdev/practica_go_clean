package aplication

import (
	"one/internal/core/users/domain"
	"strings"
)

type UserCase struct {
	repo domain.UserRepository // solo me conecto a la forma de UserRepository
}

func NewUserCase(repository domain.UserRepository) *UserCase {
	return &UserCase{repo: repository}
}

// Implemenationc del Caso de uso 1
func (uc *UserCase) ObtenerUsuarios() ([]domain.UserEntity, error) {
	// pide datos al repo
	users, err := uc.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// logica 1 => si no hay usuarios lanzamos error de negocio
	if len(users) == 0 {
		return []domain.UserEntity{}, nil // lista vacia mejor
	}

	// Logica 2 => Ocultar edad de menores de edad
	for i := range users {
		if users[i].Edad < 18 {
			users[i].Edad = 0
		}
		// Standarizar email
		users[i].Email = strings.ToLower(users[i].Email)
	}

	return users, nil
}
