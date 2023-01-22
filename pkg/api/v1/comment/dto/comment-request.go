package dto

import (
	"strings"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/entities"
)

type CommentRequestBody struct {
	PostID	int			`json:"post_id"`
	Body		string	`json:"body"`
}

func (p CommentRequestBody) ToCommentEntity() *entities.Comment {
	return &entities.Comment{
		PostID:	p.PostID,
		Body:		strings.TrimSpace(p.Body),
	}
}
