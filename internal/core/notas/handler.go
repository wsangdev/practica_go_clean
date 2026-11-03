package notas

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotasHanlder struct {
	usecase NotasUseCase // inyactamos la interface
}

func NewNotasHandler(uc NotasUseCase) *NotasHanlder {
	return &NotasHanlder{usecase: uc}
}

// enpotin para mostrar todas las notas
func (h *NotasHanlder) GetAll(c *gin.Context) {
	notas, err := h.usecase.ObtenerAllNotas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notas)
}

// enpoint para crear una nota
func (h *NotasHanlder) CreateNota(c *gin.Context) {
	var nota Nota

	// validacion del json correcto
	if err := c.ShouldBindJSON(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido: " + err.Error()})
		return
	}

	if err := h.usecase.CreateNota(&nota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nota)
}
