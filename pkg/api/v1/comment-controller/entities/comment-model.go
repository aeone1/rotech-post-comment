package models

import "time"

// A Comment Model
// With post_id to link it to posts
type Comment struct {
	ID				int				`db:"id"`
	Body			string		`db:"body"`
	PostID		int				`db:"post_id"`
	Created  	time.Time	`db:"created_at"`
}
