package entities

import "time"

// A Comment Model
// With post_id to link it to posts
type Comment struct {
	ID				int				`db:"id"`
	Body			string		`db:"body"`
	PostID		int				`db:"post_id"`
	CreatedAt time.Time	`db:"created_at"`
	DeletedAt	time.Time	`db:"deleted_at"`
}

type CommentsList []*Comment

func (p *Comment) TableName() string {
	return "comments"
}
