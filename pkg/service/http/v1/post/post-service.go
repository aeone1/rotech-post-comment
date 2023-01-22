package post

import (
	"log"
	"strconv"

	httpresponse "github.com/aeone1/rotech-post-comment/internal/protocols/http/response"
	commentRepository "github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/repository"
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
	err := c.Bind(&postReq)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	postEntity := postReq.ToPostEntity()
	err = repository.CreatePost(tx, postEntity)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	postResp := dto.CreatePostResponse(*postEntity)
	httpresponse.Json(c, 200, "created", postResp)
}

func (p *PostService) GetPostByID (c *gin.Context) {
	rawPostId := c.Param("id")
	postId, err := strconv.Atoi(rawPostId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	post, err := repository.GetPostByID(tx, postId)
	if err != nil {
		httpresponse.Err(c, err)
		return
	}
	postResp := dto.CreatePostResponse(*post)
	httpresponse.Json(c, 200, "fetchById", postResp)
}

func (p *PostService) GetPosts (c *gin.Context) {
	tx := p.db.MustBegin()
	posts, err := repository.GetPosts(tx)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	postResp := dto.CreatePostsListResponse(posts)
	httpresponse.Json(c, 200, "fetchAll", postResp)
}

func (p *PostService) GetPostsCount (c *gin.Context) {
	tx := p.db.MustBegin()
	count, err := repository.GetPostsCount(tx)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	countPostResp := dto.CreateCountResponse(*count)
	httpresponse.Json(c, 200, "CountAllPost", countPostResp)
}


func (p *PostService) UpdatePost (c *gin.Context) {
	rawPostId := c.Param("id")
	postId, err := strconv.Atoi(rawPostId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	postReq := dto.PostRequestBody{}
	err = c.Bind(&postReq)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	postEntity := postReq.ToPostEntity()
	postEntity.ID = postId
	err = repository.UpdatePost(tx, postEntity)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	postResp := dto.CreatePostResponse(*postEntity)
	httpresponse.Json(c, 200, "updated", postResp)
}

func (p *PostService) DeletePost (c *gin.Context) {
	rawPostId := c.Param("id")
	postId, err := strconv.Atoi(rawPostId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	deleteResult, deleteError := repository.DeletePost(tx, postId)
	if deleteError != nil {
		tx.Rollback()
		log.Println(deleteError)
		httpresponse.Err(c, deleteError)
		return
	}
	err = commentRepository.DeleteCommentByPostID(tx, postId)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	if deleteResult {
		httpresponse.Json(c, 200, "Post was successfully deleted", deleteResult)
		return
	}
	httpresponse.Json(c, 400, "Post was not successfully deleted", deleteResult)
}
