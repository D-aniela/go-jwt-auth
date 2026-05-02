package main

// Server creating a new instance of Fiber

import (
	"jwt-auth-api/database"
	"jwt-auth-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	database.DBconn()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
	}))

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
