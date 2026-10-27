package infra

import "one/internal/core/users/domain"

func ToDomain(m *UserModel) *domain.UserEntity {
	return &domain.UserEntity{
		ID:     m.ID,
		Nombre: m.Nombre,
		Email:  m.Email,
		Edad:   m.Edad,
	}
}

func ToDomainList(models []UserModel) []domain.UserEntity {
	entities := make([]domain.UserEntity, len(models))
	for i, m := range models {
		entities[i] = *ToDomain(&m)
	}
	return entities
}

func ToModel(e *domain.UserEntity) *UserModel {
	return &UserModel{
		ID:     e.ID,
		Nombre: e.Nombre,
		Email:  e.Email,
		Edad:   e.Edad,
	}
}
