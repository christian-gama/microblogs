package main

import (
	"log"

	"github.com/christian-gama/shared/event"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	postStore := NewPostStore()
	app := fiber.New()
	app.Use(cors.New(cors.Config{}))

	app.Post("/posts/create", func(c *fiber.Ctx) error {
		log.Println("Saving post")
		post, err := parsePost(c)
		if err != nil {
			return c.Status(400).SendString("Failed to parse post")
		}

		e := postStore.AddPost(post)
		go event.Send(e)

		return c.JSON(post)
	})

	app.Post("/events", func(c *fiber.Ctx) error {
		log.Println("Received event")
		return nil
	})

	app.Listen(":4000")
}
