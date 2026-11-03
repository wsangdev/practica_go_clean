package notas

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup, h *NotasHanlder) {
	notas := router.Group("/notas")
	{
		notas.GET("", h.GetAll)
		notas.POST("", h.CreateNota)
	}
}
