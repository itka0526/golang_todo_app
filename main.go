package main

import (
	"golang_project_2/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

const PORT = "localhost:3000"

func main() {
	var app *fiber.App = fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(PORT))
}