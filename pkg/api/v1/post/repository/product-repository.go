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

// func GetPost(tx *sqlx.Tx, post *entities.Post)
