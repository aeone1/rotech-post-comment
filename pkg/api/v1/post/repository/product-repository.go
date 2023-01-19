package repository

import (
	"fmt"
	"log"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/entities"

	"github.com/jmoiron/sqlx"
)

// Inserts Post to DB
// Returns error if error was generated
func CreatePost(tx *sqlx.Tx, post *entities.Post) error {
	rows, err := tx.NamedQuery(
		fmt.Sprintf("INSERT INTO %s (title, body) VALUES (:title, :body) RETURNING id, title, body, created_at", post.TableName()),
		&post,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next() {
		err := rows.StructScan(&post)
		if err != nil {
			return err
		}
	}
	fmt.Printf("%#v\n", post)

	return nil
}

// Fetch all posts
// TODO Add user filtering
func GetPosts(tx *sqlx.Tx) ([]*entities.Post, error) {
	posts := []*entities.Post{}
	err := tx.Select(
		&posts,
		`SELECT id, title, body, created_at 
		FROM posts
		WHERE deleted_at IS NULL
		ORDER BY id DESC`,
	)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Fetch post by ID
func GetPostByID(tx *sqlx.Tx, postId int) (*entities.Post, error) {
	post := entities.Post{}
	rows, err := tx.Queryx(
		`SELECT id, title, body, created_at 
		FROM posts
		WHERE 
			id = $1
			AND deleted_at IS NULL
		ORDER BY id DESC`,
		postId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&post)
		if err != nil {
			return nil, err
		}
	}
	return &post, nil
}
