package main

import (
	"strings"

	"github.com/christian-gama/shared/event"
)

func ModerateComment(comment *Comment) *event.Event[*Comment] {
	if strings.Contains(comment.Content, "orange") {
		comment.Status = "rejected"
	} else {
		comment.Status = "approved"
	}

	return &event.Event[*Comment]{
		Type: "CommentModerated",
		Data: comment,
	}
}
