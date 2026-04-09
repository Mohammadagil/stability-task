package main

import (
	"log"
	"stability-test-task-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Optional: Tambah logging middleware
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("[%s] %s", c.Method(), c.Path())
		return c.Next()
	})

	// Routes
	app.Get("/tasks", handlers.GetTasks)
	app.Get("/tasks/:id", handlers.GetTask)
	app.Post("/tasks", handlers.CreateTask)
	app.Put("/tasks/:id", handlers.UpdateTask)  // ROUTE BARU
	app.Delete("/tasks/:id", handlers.DeleteTask)

	// Start server
	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
