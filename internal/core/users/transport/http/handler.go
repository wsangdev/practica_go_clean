package http

import (
	"net/http"
	"one/internal/core/users/domain"

	"github.com/gin-gonic/gin"
)

type UserCase interface {
	ObtenerUsuarios() ([]domain.UserEntity, error)
}

type UserHandler struct {
	uc UserCase
}

// funcion constructro
func NewUserHandler(uc UserCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// funcion para Mostrar todos los usuarios
func (h *UserHandler) GetAllUsers(c *gin.Context) {

	users, err := h.uc.ObtenerUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := make([]UserResponse, 0)
	for _, u := range users {
		response = append(response, UserResponse{
			ID:     u.ID,
			Nombre: u.Nombre,
			Email:  u.Email,
			Edad:   u.Edad,
		})
	}

	c.JSON(http.StatusOK, response)
}
