package service

import (

	"github.com/gin-gonic/gin"
)

type PostService interface {
	GetPostByID(c *gin.Context)
	GetPosts(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}
