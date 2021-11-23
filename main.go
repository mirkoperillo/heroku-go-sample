package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
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
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Println("env PORT not configured, set the default 3000")
		serverPort = "3000"
	}

	server := fiber.New()
	server.Get("/api/hello", hello)
	server.Static("/", "./public")
	server.Listen(fmt.Sprintf(":%s", serverPort))
}
