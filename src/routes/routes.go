package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateReqBody struct {
	Message string `json:"message"`
}

type UpdateReqBody struct {
	Id uuid.UUID `json:"id"`
	Status string `json:"status"`
}

type DeleteReqBody struct {
	Id uuid.UUID `json:"id"`
}

type Task struct {
	Id uuid.UUID `json:"id"`
	TaskName string `json:"taskName"`
	Status string `json:"status"`
}

func Routes(app *fiber.App) {
	app.Static("/", "src/frontend/dist")

	app.All("/api/tasks", func(c *fiber.Ctx) error {
		method := c.Method()
		switch method {
		case "POST":
			return Create(c)
		case "GET":
			return Read(c)
		case "PUT":
			return Update(c)
		case "DELETE":
			return Delete(c)
		default:
			return c.SendString("Route is not being handled")
		}
	})
}
