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
		return c.JSON(postStore.GetAllPosts())
	})

	app.Post("/events", func(c *fiber.Ctx) error {
		event := &event.Event[map[string]any]{}

		if err := c.BodyParser(event); err != nil {
			return err
		}

		if event.Type == "PostCreated" {
			post := parsePost(event.Data)
			postStore.AddPost(post)
		}

		if event.Type == "CommentCreated" {
			comment := parseComment(event.Data)
			postStore.AddCommentToPost(comment)
		}

		if event.Type == "CommentUpdated" {
			comment := parseComment(event.Data)
			postStore.UpdateCommentInPost(comment)
		}

		return nil
	})

	app.Listen(":4002")
}
