package main

import (
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = []Task{
	{ID: 1, Title: "Learn Go", Status: "pending"},
}

func main() {
	app := fiber.New()

	// Routes
	app.Get("/tasks", getTasks)
	app.Post("/tasks", addTask)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)

	app.Listen(":3000")
}

func getTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}

func addTask(c *fiber.Ctx) error {
	var newTask Task
	if err := c.BodyParser(&newTask); err != nil {
		return c.Status(400).SendString("Invalid input")
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	return c.JSON(newTask)
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, task := range tasks {
		if id == string(rune(task.ID)) {
			if err := c.BodyParser(&tasks[i]); err != nil {
				return c.Status(400).SendString("Invalid input")
			}
			return c.JSON(tasks[i])
		}
	}
	return c.Status(404).SendString("Task not found")
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, task := range tasks {
		if id == string(rune(task.ID)) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.SendString("Task deleted")
		}
	}
	return c.Status(404).SendString("Task not found")
}

