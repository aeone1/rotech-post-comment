package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectToDB() {
	var err error
	host := "localhost"
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
			log.Fatalln(err)
	}
}
