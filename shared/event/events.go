package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Event[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

func Parse[T any](c *fiber.Ctx) (*Event[T], error) {
	event := &Event[T]{}
	if err := c.BodyParser(event); err != nil {
		return nil, err
	}
	return event, nil
}

func Create[T any](eventType string, data T) *Event[T] {
	return &Event[T]{
		Type: eventType,
		Data: data,
	}
}

func Send[T any](event *Event[T]) error {
	marshalledEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println("Sending event to event bus")
	_, err = http.Post("http://eventbus-clusterip-srv:4005/events", "application/json", bytes.NewReader(marshalledEvent))
	if err != nil {
		log.Println("EventBus is down", err)
		return err
	}

	return nil
}
