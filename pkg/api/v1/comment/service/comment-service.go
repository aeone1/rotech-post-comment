package service

import (

	"github.com/gin-gonic/gin"
)

type CommentService interface {
	GetCommentByID(c *gin.Context)
	GetComments(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}
