package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Correct struct declaration
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!" + " " + "Greetings, Universe!")
	app := fiber.New()

	// Correct variable name and initialization
	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"msg": "hello world",
		})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// Correct method name for parsing
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Validate the todo body
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		// Assigning the correct ID
		todo.ID = len(todos) + 1

		// Append the todo to the slice
		todos = append(todos, *todo)

		// Return the newly created todo
		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))
}
