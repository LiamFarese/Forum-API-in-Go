package comment

import (
	"database/sql"

	"github.com/jackc/pgx/pgtype"
)

type Comment struct {
	ID            int64              `db:"id" json:"id"`
	Body          string             `db:"body" json:"body"`
	UserID        int64              `db:"user_id" json:"user_id"`
	PostID        int64              `db:"post_id" json:"post_id"`
	ParentComment sql.NullInt64      `db:"parent_comment" json:"parent_comment"`
	CreatedAt     pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type CommentWithUsername struct {
	ID            int64              `db:"id" json:"id"`
	Username      string             `db:"username" json:"username"`
	Body          string             `db:"body" json:"body"`
	UserID        int64              `db:"user_id" json:"user_id"`
	PostID        int64              `db:"post_id" json:"post_id"`
	ParentComment sql.NullInt64      `db:"parent_comment" json:"parent_comment"`
	CreatedAt     pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type CreateCommentParams struct {
	Body   string `json:"title"`
	UserID int64  `json:"user_id"`
	PostID int64  `json:"post_id"`
}

type CreateReplyParams struct {
	Body          string        `json:"title"`
	UserID        int64         `json:"user_id"`
	PostID        int64         `json:"post_id"`
	ParentComment sql.NullInt64 `json:"parent_comment"`
}
