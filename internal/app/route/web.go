package route

import (
	"github.com/Noringotq/go-crud/internal/pkg/task"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/todos", task.Todos)
	api.Get("/todo/:id", task.Todo)
	api.Post("/todo", task.TodoStore)
	api.Put("/todo/:id", task.TodoUpdate)
	api.Delete("/todo/:id", task.TodoDelete)
}
