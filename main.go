package main

import (
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Greets string `json:"greets"`
}

func hello(c *fiber.Ctx) error {
	payload := Payload{"Gophers"}
	c.JSON(payload)
	return nil
}

func main() {
	server := fiber.New()
	server.Get("/api/hello", hello)
	server.Static("/", "./public")
	server.Listen(":3000")
}
