package comment

import (
	"errors"
	"log"
	"strconv"

	httpresponse "github.com/aeone1/rotech-post-comment/internal/protocols/http/response"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/dto"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/repository"
	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/service"
	postRepository "github.com/aeone1/rotech-post-comment/pkg/api/v1/post/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type CommentService struct {
	db *sqlx.DB
}

func NewCommentService(db *sqlx.DB) service.CommentService {
	return &CommentService{ db: db}
}

func (p *CommentService) CreateComment (c *gin.Context) {
	commentReq := dto.CommentRequestBody{}
	err := c.Bind(&commentReq)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	post, getPostErr := postRepository.GetPostByID(tx, commentReq.PostID)
	if post.ID == 0 {
		log.Println(err)
		httpresponse.Err(c, errors.New("Post was not found"))
		return
	}
	if getPostErr != nil {
		log.Println(getPostErr)
		httpresponse.Err(c, getPostErr)
		return
	}
	commentEntity := commentReq.ToCommentEntity()
	err = repository.CreateComment(tx, commentEntity)
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
	commentResp := dto.CreateCommentResponse(*commentEntity)
	httpresponse.Json(c, 200, "created", commentResp)
}

func (p *CommentService) GetCommentByID (c *gin.Context) {
	rawCommentId := c.Param("id")
	commentId, err := strconv.Atoi(rawCommentId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	comment, err := repository.GetCommentByID(tx, commentId)
	if err != nil {
		httpresponse.Err(c, err)
		return
	}
	commentResp := dto.CreateCommentResponse(*comment)
	httpresponse.Json(c, 200, "fetchById", commentResp)
}

func (p *CommentService) GetComments (c *gin.Context) {
	tx := p.db.MustBegin()
	comments, err := repository.GetAllComments(tx)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	commentResp := dto.CreateCommentsListResponse(comments)
	httpresponse.Json(c, 200, "fetchAll", commentResp)
}

func (p *CommentService) UpdateComment (c *gin.Context) {
	rawCommentId := c.Param("id")
	commentId, err := strconv.Atoi(rawCommentId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	commentReq := dto.CommentRequestBody{}
	err = c.Bind(&commentReq)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	commentEntity := commentReq.ToCommentEntity()
	commentEntity.ID = commentId
	err = repository.UpdateComment(tx, commentEntity)
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
	commentResp := dto.CreateCommentResponse(*commentEntity)
	httpresponse.Json(c, 200, "updated", commentResp)
}

func (p *CommentService) DeleteComment (c *gin.Context) {
	rawCommentId := c.Param("id")
	commentId, err := strconv.Atoi(rawCommentId)
	if err != nil {
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	tx := p.db.MustBegin()
	deleteResult, deleteError := repository.DeleteCommentByID(tx, commentId)
	if deleteError != nil {
		tx.Rollback()
		log.Println(deleteError)
		httpresponse.Err(c, deleteError)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		httpresponse.Err(c, err)
		return
	}
	if deleteResult {
		httpresponse.Json(c, 200, "Comment was successfully deleted", deleteResult)
		return
	}
	httpresponse.Json(c, 400, "Comment was not successfully deleted", deleteResult)
}
