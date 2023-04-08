package main

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var events = []any{}

func main() {
	app := fiber.New()

	app.Post("/events", func(c *fiber.Ctx) error {
		events = append(events, c.Body())

		go http.Post("http://localhost:4000/events", "application/json", bytes.NewReader(c.Body()))
		go http.Post("http://localhost:4001/events", "application/json", bytes.NewReader(c.Body()))
		go http.Post("http://localhost:4002/events", "application/json", bytes.NewReader(c.Body()))
		go http.Post("http://localhost:4003/events", "application/json", bytes.NewReader(c.Body()))

		return nil
	})

	app.Get("/events", func(c *fiber.Ctx) error {
		return c.JSON(events)
	})

	app.Listen(":4005")
}
