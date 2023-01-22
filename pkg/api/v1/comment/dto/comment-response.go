package dto

import (
	"time"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/entities"
)

type CommentResponse struct {
	ID				int				`json:"id"`
	PostID		int				`json:"post_id"`
	Body			string		`json:"body"`
	CreatedAt time.Time	`json:"created_at"`
}

func CreateCommentResponse(comment entities.Comment) CommentResponse {
	return CommentResponse{
		ID:					comment.ID,
		PostID: 		comment.PostID,
		Body: 			comment.Body,
		CreatedAt: 	comment.CreatedAt,
	}
}

type CommentsListReponse []*CommentResponse

func CreateCommentsListResponse(comments entities.CommentsList) CommentsListReponse {
	commentsResp := make(CommentsListReponse, 0, len(comments))
	for _, c := range comments {
		comment := CreateCommentResponse(*c)
		commentsResp = append(commentsResp, &comment)
	}
	return commentsResp
}
