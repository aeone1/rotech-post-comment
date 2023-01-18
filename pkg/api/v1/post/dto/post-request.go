package dto

import (
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/entities"
)

type PostRequestBody struct {
	Title	string	`json:"title"`
	Body	string	`json:"body"`
}

type PostRequestParams struct {
	ID int
}

func (p PostRequestBody) ToPostEntities() *entities.Post {
	return &entities.Post{
		Title:        p.Title,
		Body: 				p.Body,
	}
}
