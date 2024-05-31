package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

func main(){
	fmt.Println("Hello world")
	app := fiber.New()

	log.Fatal(app.Listen(":4000"))
	
}
