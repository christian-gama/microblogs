package main

import (
	"github.com/christian-gama/event-bus/pkg/event"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	postStore := NewPostStore()
	app := fiber.New()
	app.Use(cors.New(cors.Config{}))

	app.Get("/posts", func(c *fiber.Ctx) error {
		posts := postStore.GetPosts()
		return c.JSON(posts)
	})

	app.Post("/posts", func(c *fiber.Ctx) error {
		post, err := parsePost(c)
		if err != nil {
			return c.Status(400).SendString("Failed to parse post")
		}

		e := postStore.AddPost(post)
		go event.Send(e)

		return c.JSON(post)
	})

	app.Listen(":4000")
}
