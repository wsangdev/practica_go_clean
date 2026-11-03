package notas

import "gorm.io/gorm"

type NotasRepository interface {
	GetAll() ([]Nota, error)
	Create(nota *Nota) error
}

type NotasGorm struct {
	db *gorm.DB
}

func NewNotasGorm(database *gorm.DB) NotasRepository {
	return &NotasGorm{db: database}
}

// Repositorio Encontrar todas las notas
func (r *NotasGorm) GetAll() ([]Nota, error) {
	var notas []Nota
	result := r.db.Find(&notas)

	if result.Error != nil {
		return nil, result.Error
	}

	return notas, nil
}

// Repo para Crear ntoas
func (r *NotasGorm) Create(nota *Nota) error {

	return r.db.Create(nota).Error
}
