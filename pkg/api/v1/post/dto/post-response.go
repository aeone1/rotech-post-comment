package dto

import (
	"time"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/entities"
)

type PostResponse struct {
	ID				int				`json:"id"`
	Title			string		`json:"title"`
	Body			string		`json:"body"`
	CreatedAt time.Time	`json:"created_at"`
}

func CreatePostResponse(post entities.Post) PostResponse {
	return PostResponse{
		ID: 				post.ID,
		Title: 			post.Title,
		Body: 			post.Body,
		CreatedAt: 	post.CreatedAt,
	}
}

type PostsListReponse []*PostResponse

func CreatePostsListResponse(posts entities.PostsList) PostsListReponse {
	postsResp := make(PostsListReponse, 0, len(posts))
	for _, p := range posts {
		post := CreatePostResponse(*p)
		postsResp = append(postsResp, &post)
	}
	return postsResp
}

type CountResponse struct {
	Count int `json:"count"`
}

func CreateCountResponse(countPosts int) CountResponse {
	return CountResponse{
		Count: countPosts,
	}
}
