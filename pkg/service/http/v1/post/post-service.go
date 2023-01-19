package post

import (
	"fmt"
	"strconv"

	httpresponse "github.com/aeone1/rotech-post-comment/internal/protocols/http/response"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/dto"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/repository"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/service"

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
	rawPostId := c.Param("id")
	postId, err := strconv.Atoi(rawPostId)
	if err != nil {
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	post, err := repository.GetPostByID(tx, postId)
	if err != nil {
		httpresponse.Err(c, err)
	}
	postResp := dto.CreatePostResponse(*post)
	httpresponse.Json(c, 200, "fetchById", postResp)
}

func (p *PostService) GetPosts (c *gin.Context) {
	tx := p.db.MustBegin()
	posts, err := repository.GetPosts(tx)
	if err != nil {
		httpresponse.Err(c, err)
	}
	postResp := dto.CreatePostsListResponse(posts)
	httpresponse.Json(c, 200, "fetchAll", postResp)
}

func (p *PostService) UpdatePost (c *gin.Context) {

	c.String(200, "pong")
}

func (p *PostService) DeletePost (c *gin.Context) {
	c.String(200, "pong")
}
