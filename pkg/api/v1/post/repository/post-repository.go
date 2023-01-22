package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/post/entities"

	"github.com/jmoiron/sqlx"
)

// Inserts Post to DB
// Returns error if error was generated
func CreatePost(tx *sqlx.Tx, post *entities.Post) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (title, body) VALUES (:title, :body) RETURNING id, title, body, created_at",
		post.TableName(),
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &post)
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

	return nil
}

// Fetch all posts
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

// Count all posts
func GetPostsCount(tx *sqlx.Tx) (*int, error) {
	var postsCount int
	rows, err := tx.Queryx(
		`SELECT count(id) as postsCount
		FROM posts
		WHERE deleted_at IS NULL`,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&postsCount)
		if err != nil {
			return nil, err
		}
	}
	return &postsCount, nil
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

// Updates Post in DB by ID
// Returns error if error was generated
func UpdatePost(tx *sqlx.Tx, post *entities.Post) error {
	queryParams := ""
	if post.Body != "" {
		if queryParams != "" {
			queryParams += ", "
		}
		queryParams += "body =:body"
	}
	if post.Title != "" {
		if queryParams != "" {
			queryParams += ", "
		}
		queryParams += "title =:title"
	}
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id =:id AND deleted_at IS NULL RETURNING id, title, body, created_at",
		post.TableName(),
		queryParams,
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &post)
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

	return nil
}

// Inserts Post to DB
// Returns error if error was generated
func DeletePost(tx *sqlx.Tx, postId int) (bool, error) {
	post := entities.Post{}
	post.DeletedAt = time.Now()
	query := fmt.Sprintf(
		"UPDATE %s SET deleted_at = :deleted_at WHERE id = %d AND deleted_at IS NULL RETURNING id",
		post.TableName(),
		postId,
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &post)
	if err != nil {
		log.Println(err)
		return false, err
	}
	for rows.Next() {
		err := rows.StructScan(&post)
		if err != nil {
			return false, err
		}
	}

	return post.ID == postId, nil
}