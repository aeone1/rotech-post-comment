package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aeone1/rotech-post-comment/pkg/api/v1/comment/entities"

	"github.com/jmoiron/sqlx"
)

// Inserts Comment to DB
// Returns error if error was generated
func CreateComment(tx *sqlx.Tx, comment *entities.Comment) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (post_id, body) VALUES (:post_id, :body) RETURNING id, post_id, body, created_at",
		comment.TableName(),
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &comment)
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next() {
		err := rows.StructScan(&comment)
		if err != nil {
			return err
		}
	}

	return nil
}

// Fetch all comments of a post
func GetPostComments(tx *sqlx.Tx, postId int ) ([]*entities.Comment, error) {
	comments := []*entities.Comment{}
	err := tx.Select(
		&comments,
		fmt.Sprintf(
			`SELECT id, post_id, body, created_at 
			FROM comments
			WHERE deleted_at IS NULL
			AND post_id = %d
			ORDER BY id DESC`,
			postId,
		),
	)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Fetch all comments
func GetAllComments(tx *sqlx.Tx) ([]*entities.Comment, error) {
	comments := []*entities.Comment{}
	err := tx.Select(
		&comments,
		fmt.Sprintf(
			`SELECT id, post_id, body, created_at 
			FROM comments
			WHERE deleted_at IS NULL
			ORDER BY id DESC`,
		),
	)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// Fetch comment by ID
func GetCommentByID(tx *sqlx.Tx, commentId int) (*entities.Comment, error) {
	comment := entities.Comment{}
	rows, err := tx.Queryx(
		`SELECT id, body, created_at 
		FROM comments
		WHERE 
			id = $1
			AND deleted_at IS NULL
		ORDER BY id DESC`,
		commentId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
	}
	return &comment, nil
}

// Updates Comment in DB by ID
// Returns error if error was generated
func UpdateComment(tx *sqlx.Tx, comment *entities.Comment) error {
	if comment.Body == "" {
		return errors.New("Empty update Body")
	}
	query := fmt.Sprintf(
		"UPDATE %s SET body =:body WHERE id =:id AND deleted_at IS NULL RETURNING id, post_id, body, created_at",
		comment.TableName(),
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &comment)
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next() {
		err := rows.StructScan(&comment)
		if err != nil {
			return err
		}
	}

	return nil
}

// Soft Delete Comment in DB by CommentID
// Returns error if error was generated
func DeleteCommentByID(tx *sqlx.Tx, commentId int) (bool, error) {
	comment := entities.Comment{}
	comment.DeletedAt = time.Now()
	query := fmt.Sprintf(
		"UPDATE %s SET deleted_at = :deleted_at WHERE id = %d AND deleted_at IS NULL RETURNING id",
		comment.TableName(),
		commentId,
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &comment)
	if err != nil {
		log.Println(err)
		return false, err
	}
	for rows.Next() {
		err := rows.StructScan(&comment)
		if err != nil {
			return false, err
		}
	}

	return comment.ID == commentId, nil
}

// Soft Delete Comment in DB by PostID
// Returns error if error was generated
func DeleteCommentByPostID(tx *sqlx.Tx, postId int) error {
	comment := entities.Comment{}
	comment.DeletedAt = time.Now()
	query := fmt.Sprintf(
		"UPDATE %s SET deleted_at = :deleted_at WHERE post_id = %d AND deleted_at IS NULL RETURNING id",
		comment.TableName(),
		postId,
	)
	log.Println(query)
	rows, err := tx.NamedQuery(query, &comment)
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next() {
		err := rows.StructScan(&comment)
		if err != nil {
			return err
		}
	}

	return nil
}