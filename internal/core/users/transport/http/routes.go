package http

import "github.com/gin-gonic/gin"

// RegisterRoutes define las rutas de este módulo específico
func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("", h.GetAllUsers)
		// users.POST("", h.CreateUser) // Aquí agregarás el futuro
	}
}
