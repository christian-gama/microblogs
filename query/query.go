package main

import (
	"errors"
)

type Post struct {
	ID       string     `json:"id"`
	Title    string     `json:"title"`
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
	PostID  string `json:"postId"`
}

type PostStore struct {
	posts map[string]*Post
}

func NewPostStore() *PostStore {
	return &PostStore{
		posts: make(map[string]*Post),
	}
}

func (ps *PostStore) GetAllPosts() map[string]*Post {
	return ps.posts
}

func (ps *PostStore) AddPost(post *Post) {
	ps.posts[post.ID] = post
}

func (ps *PostStore) AddCommentToPost(comment *Comment) error {
	post, ok := ps.posts[comment.PostID]
	if !ok {
		return errors.New("post not found")
	}

	post.Comments = append(post.Comments, comment)
	ps.posts[post.ID] = post

	return nil
}

func (ps *PostStore) UpdateCommentInPost(comment *Comment) error {
	post, ok := ps.posts[comment.PostID]
	if !ok {
		return errors.New("post not found")
	}

	for i, c := range post.Comments {
		if c.ID == comment.ID {
			post.Comments[i] = comment
			ps.posts[post.ID] = post
			break
		}
	}

	return nil
}

func parsePost(data interface{}) *Post {
	dataMap := data.(map[string]interface{})
	return &Post{
		ID:    dataMap["id"].(string),
		Title: dataMap["title"].(string),
	}
}

func parseComment(data interface{}) *Comment {
	dataMap := data.(map[string]interface{})
	return &Comment{
		ID:      dataMap["id"].(string),
		Content: dataMap["content"].(string),
		Status:  dataMap["status"].(string),
		PostID:  dataMap["postId"].(string),
	}
}
