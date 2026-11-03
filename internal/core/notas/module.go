package notas

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitNotasModule(router *gin.RouterGroup, db *gorm.DB) {

	repo := NewNotasGorm(db)
	usecase := NewNotasCase(repo)
	handler := NewNotasHandler(usecase)

	RegisterRouter(router, handler)
}
