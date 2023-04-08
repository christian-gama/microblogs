package main

import (
	"github.com/christian-gama/shared/event"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	commentStore := NewCommentStore()
	app := fiber.New()
	app.Use(cors.New(cors.Config{}))

	app.Get("/posts/:id/comments", func(c *fiber.Ctx) error {
		comments, ok := commentStore.GetComments(c.Params("id"))
		if !ok {
			return c.SendStatus(404)
		}
		return c.JSON(comments)
	})

	app.Post("/posts/:id/comments", func(c *fiber.Ctx) error {
		comment, err := parseComment(c)
		if err != nil {
			return err
		}

		e := commentStore.AddComment(comment.PostID, comment)
		go event.Send(e)
		return c.JSON(comment)
	})

	app.Post("/events", func(c *fiber.Ctx) error {
		e, err := event.Parse[*Comment](c)
		if err != nil {
			return err
		}

		switch e.Type {
		case "CommentModerated":
			e := commentStore.UpdateCommentStatus(e.Data.PostID, e.Data.ID, e.Data.Status)
			go event.Send(e)
		}

		return nil
	})

	app.Listen(":4001")
}
