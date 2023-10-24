package post

import "github.com/jackc/pgx/pgtype"

type Post struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	// Content of the listing
	Body      string             `db:"body" json:"body"`
	UserID    int64              ` db:"user_id" json:"user_id"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type PostWithUsername struct {
	ID       int64  `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Username string `db:"username" json:"username"`
	// Content of the listing
	Body      string             `db:"body" json:"body"`
	UserID    int64              ` db:"user_id" json:"user_id"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type CreatePostParams struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"user_id"`
}
