package entities

import "time"

// A Post Model/Entity
// I haven't included the post creator id
// as I didn't implement user creation and
// authentification/authorization
type Post struct {
	ID				int				`db:"id"`
	Title			string		`db:"title"`
	Body			string		`db:"body"`
	CreatedAt time.Time	`db:"created_at"`
	DeletedAt time.Time	`db:"deleted_at"`
}

type PostsList []*Post

func (p *Post) TableName() string {
	return "posts"
}
