package notas

import "gorm.io/gorm"

type NotasRepository interface {
	GetAll() ([]Nota, error)
	Create(nota *Nota) error
}

type NotasGorm struct {
	db *gorm.DB 
}

