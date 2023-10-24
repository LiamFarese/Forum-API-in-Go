package listing

import "github.com/jackc/pgx/pgtype"

type Listing struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	// Content of the listing
	Body      string             `db:"body" json:"body"`
	UserID    int64              ` db:"user_id" json:"user_id"`
	IsSettled bool               ` db:"is_settled" json:"is_settled"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type ListingWithUsername struct {
	ID       int64  `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Username string `db:"username" json:"username"`
	// Content of the listing
	Body      string             `db:"body" json:"body"`
	UserID    int64              ` db:"user_id" json:"user_id"`
	IsSettled bool               ` db:"is_settled" json:"is_settled"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type CreateListingParams struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int64  `json:"user_id"`
}
