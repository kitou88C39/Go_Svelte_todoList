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

	todos := []Todo{}

	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg":"hello world"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})
}
todo.ID = len(todos) + 1
todos = append(todos, *todo)

var x int = 5
var p *int = &x

fmt.Println(p)
fmt.Println(*p)
return c.Status(201).JSON((todo))


})

	log.Fatal(app.Listen(":4000"))
}
