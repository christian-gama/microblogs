package main

import (
	"log"

	"github.com/christian-gama/shared/event"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	PostID  string `json:"postId"`
	Status  string `json:"status"`
}

func main() {
	app := fiber.New()

	app.Post("/events", func(c *fiber.Ctx) error {
		log.Println("Received event")

		e, err := event.Parse[*Comment](c)
		if err != nil {
			return err
		}

		if e.Type == "CommentCreated" {
			e := ModerateComment(e.Data)
			go event.Send(e)
		}

		return nil
	})

	app.Listen(":4003")
}
