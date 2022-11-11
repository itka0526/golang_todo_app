package main

import (
	"golang_project_2/src/routes"
	"log"
	"os"
	"os/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var PORT string = ":"+ os.Getenv("PORT")

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if user.Username == "itgelt" {
		PORT = "localhost:3000"
	}

	var app *fiber.App = fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(PORT))
}