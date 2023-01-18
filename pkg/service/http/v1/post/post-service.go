package post

import (
	"fmt"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/dto"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/repository"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/service"
	httpresponse "github.com/aeone1/rotech-post-comment/internal/protocols/http/response"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type PostService struct {
	db *sqlx.DB
}

func NewPostService(db *sqlx.DB) service.PostService {
	return &PostService{ db: db}
}

func (p *PostService) CreatePost (c *gin.Context) {
	postReq := dto.PostRequestBody{}
	c.Bind(&postReq)
	fmt.Println(postReq)
	tx := p.db.MustBegin()
	postEntity := postReq.ToPostEntities()
	err := repository.CreatePost(tx, postEntity)
	if err != nil {
		httpresponse.Err(c, err)
	}
	postResp := dto.CreatePostResponse(*postEntity)
	httpresponse.Json(c, 200, "created", postResp)
}

func (p *PostService) GetPostByID (c *gin.Context) {
	c.String(200, "pong")
}

func (p *PostService) GetPosts (c *gin.Context) {
	c.String(200, "pong")
}

func (p *PostService) UpdatePost (c *gin.Context) {
	c.String(200, "pong")
}

func (p *PostService) DeletePost (c *gin.Context) {
	c.String(200, "pong")
}
