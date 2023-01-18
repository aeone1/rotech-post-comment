package main

import (
	"github.com/aeone1/rotech-post-comment/initializers"
)

// DB Schema of the App
// Postgres dialect
var schema = `
CREATE TABLE IF NOT EXISTS posts (
	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
	title varchar(255),
	body text,
	created_at timestamp(6) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp(6) without time zone,
	CONSTRAINT posts_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS comments (
	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
	body text,
	post_id bigint NOT NULL,
	created_at timestamp(6) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp(6) without time zone,
	CONSTRAINT comments_pkey PRIMARY KEY (id),
	CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id)
		REFERENCES posts (id) MATCH SIMPLE
		ON UPDATE CASCADE
		ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS posts_title_idx
    ON posts USING btree
    (title ASC NULLS LAST)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS comments_post_id_idx
    ON comments USING btree
    (post_id ASC NULLS LAST)
    TABLESPACE pg_default;
`

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	initializers.DB.MustExec(schema)
}