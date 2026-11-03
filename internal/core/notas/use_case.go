package notas

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type NotasUseCase interface {
	ObtenerAllNotas() ([]Nota, error)
	CreateNota(nota *Nota) error
}

type NotasCase struct {
	repo NotasRepository
}

func NewNotasCase(repository NotasRepository) *NotasCase {
	return &NotasCase{repo: repository}
}

func (uc *NotasCase) ObtenerAllNotas() ([]Nota, error) {
	return uc.repo.GetAll()
}

func (uc *NotasCase) CreateNota(note *Nota) error {

	if note.Titulo == "" {
		return errors.New("el titulo de la nota es obligatorio")
	}

	note.ID = uuid.New().String()
	note.CreatedAt = time.Now()

	return uc.repo.Create(note)
}
