package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conect() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se encontró el archivo .env, usando variables del sistema")
	}

	// 2. Leer la URL de conexión de Neon
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("la variable DATABASE_URL no está configurada en el .env")
	}

	// 3. Conectar a Neon (Postgres)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a Neon: %w", err)
	}

	log.Println("✅ Conexión exitosa a la base de datos de Neon")
	return db, nil
}
