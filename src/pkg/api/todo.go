package api

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Go"},
	{ID: 2, Title: "Build a Fiber API"},
}

func GetAllTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var newTodo Todo
	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	return c.JSON(newTodo)
}

func SetupTodoRoutes(group fiber.Router) {
	group.Get("/", GetAllTodos)
	group.Post("/", CreateTodo)
	// Add similar routes for updating and deleting todos.
}

// Similarly, implement functions for updating and deleting todos.
