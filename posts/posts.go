package main

import (
	"sync"

	"github.com/christian-gama/shared/event"
	"github.com/christian-gama/shared/utils"
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type PostStore struct {
	posts map[string]*Post
	mu    sync.Mutex
}

func NewPostStore() *PostStore {
	return &PostStore{
		posts: make(map[string]*Post),
	}
}

func (ps *PostStore) GetPosts() map[string]*Post {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	return ps.posts
}

func (ps *PostStore) AddPost(post *Post) *event.Event[*Post] {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.posts[post.ID] = post

	return event.Create("PostCreated", post)
}

func parsePost(c *fiber.Ctx) (*Post, error) {
	post := &Post{}
	if err := c.BodyParser(post); err != nil {
		return nil, err
	}
	post.ID = utils.GenerateID()
	return post, nil
}
