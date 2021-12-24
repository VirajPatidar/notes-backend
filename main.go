package main

import (
	"github.com/VirajPatidar/NotesBackend/database"
	"github.com/VirajPatidar/NotesBackend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"fmt"
)

func main() {
	database.Connect()

	app := fiber.New()

	port := os.Getenv("PORT")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://go-notes.netlify.app",
		AllowMethods: "GET, OPTIONS, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders:  "Origin, Content-Type, Set-Cookie, Content-Length, Accept, Cookie, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%s", port))
}
