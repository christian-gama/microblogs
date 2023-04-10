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

	app.Get("/posts", func(c *fiber.Ctx) error {
		log.Println("Retrieving posts")

		return c.JSON(postStore.GetAllPosts())
	})

	app.Post("/events", func(c *fiber.Ctx) error {
		log.Println("Received event")
		event := &event.Event[map[string]any]{}

		if err := c.BodyParser(event); err != nil {
			return err
		}

		HandleEvent(event, postStore)

		return nil
	})

	RetrieveEvents()

	app.Listen(":4002")
}
