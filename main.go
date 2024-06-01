package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){
	fmt.Println("Hello world")
	app := fiber.New()

	todos := []todo{}

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})
}

})

	log.Fatal(app.Listen(":4000"))
}
