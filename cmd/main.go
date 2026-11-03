package main

import (
	"log"
	"one/internal/app"
	"one/internal/shared/database"
)

func main() {

	// Inica base de datos
	db, err := database.Conect()
	if err != nil {
		log.Fatal("Error conectando a la BD: ", err)
	}
	// Crea la app
	myApp := app.NewApp(db)

	// levanta server
	myApp.Router.Run(":8082")
}
