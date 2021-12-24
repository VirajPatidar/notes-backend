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
		AllowOrigins: "http://localhost:3000/",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(fmt.Sprintf(":%s", port))
}
