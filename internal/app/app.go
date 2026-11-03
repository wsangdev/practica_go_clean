package app

import (
	"one/internal/core/notas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	route := gin.Default() // Definicion de la ruta

	// 1 AUTOmigraciones
	db.AutoMigrate(&notas.Nota{})

	// 2 Punto de entra de las rutas
	api := route.Group("/api/v1")

	// 3 Inicializamos modulo de notas
	notas.InitNotasModule(api, db)

	return &App{
		DB:     db,
		Router: route,
	}
}
