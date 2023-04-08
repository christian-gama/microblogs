package main

import (
	"sync"

	"github.com/christian-gama/shared/event"
	"github.com/christian-gama/shared/utils"
	"github.com/gofiber/fiber/v2"
)

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	PostID  string `json:"postId"`
	Status  string `json:"status"`
}

type CommentStore struct {
	comments map[string][]*Comment
	mu       sync.Mutex
}

func NewCommentStore() *CommentStore {
	return &CommentStore{
		comments: make(map[string][]*Comment),
	}
}

func (cs *CommentStore) GetComments(postID string) ([]*Comment, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	comments, ok := cs.comments[postID]
	return comments, ok
}

func (cs *CommentStore) AddComment(postID string, comment *Comment) *event.Event[*Comment] {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.comments[postID] = append(cs.comments[postID], comment)

	return event.Create("CommentCreated", comment)
}

func (cs *CommentStore) UpdateCommentStatus(postID, commentID, status string) *event.Event[*Comment] {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	comments, ok := cs.comments[postID]
	if !ok {
		return nil
	}

	for _, comment := range comments {
		if comment.ID == commentID {
			comment.Status = status
			return event.Create("CommentUpdated", comment)
		}
	}

	return nil
}

func parseComment(c *fiber.Ctx) (*Comment, error) {
	comment := &Comment{}
	if err := c.BodyParser(comment); err != nil {
		return nil, err
	}
	comment.ID = utils.GenerateID()
	comment.PostID = c.Params("id")
	comment.Status = "pending"
	return comment, nil
}
