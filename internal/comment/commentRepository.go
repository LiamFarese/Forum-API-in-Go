package comment

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	CreateComment(comment Comment) (CommentWithUsername, error)
	Reply(comment Comment) (CommentWithUsername, error)
	GetCommentsFromPost(postId int64) ([]CommentWithUsername, error)
}

type commentQueries struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentQueries{db: db}
}

func (q *commentQueries) CreateComment(comment Comment) (CommentWithUsername, error) {
	query := "INSERT INTO comments (body, user_id, post_id) VALUES ($1, $2, $3) RETURNING *"
	newComment := CommentWithUsername{}

	err := q.db.Get(&newComment, query, comment.Body, comment.UserID, comment.PostID)
	if err != nil {
		return CommentWithUsername{}, err
	}

	fetchUsername := "SELECT comments.*, users.username FROM comments JOIN users ON comments.user_id = users.id WHERE comments.id = $1"
	err = q.db.Get(&newComment, fetchUsername, newComment.ID)
	if err != nil {
		return CommentWithUsername{}, fmt.Errorf("could not retrieve new comment %w", err)
	}

	return newComment, nil
}

func (q *commentQueries) Reply(comment Comment) (CommentWithUsername, error) {
	query := "INSERT INTO comments (body, user_id, post_id, parent_comment) VALUES ($1, $2, $3) RETURNING *"
	newComment := CommentWithUsername{}

	err := q.db.Get(&newComment, query, comment.Body, comment.UserID, comment.PostID)
	if err != nil {
		return CommentWithUsername{}, err
	}

	fetchUsername := "SELECT comments.*, users.username FROM comments JOIN users ON comments.user_id = users.id WHERE comments.id = $1"
	err = q.db.Get(&newComment, fetchUsername, newComment.ID)
	if err != nil {
		return CommentWithUsername{}, fmt.Errorf("could not retrieve new comment %w", err)
	}

	return newComment, nil
}

func (q *commentQueries) GetCommentsFromPost(postId int64) ([]CommentWithUsername, error) {
	query := "SELECT comments.*, users.username FROM comments JOIN users ON comments.user_id = users.id WHERE comments.post_id = $1"
	comments := make([]CommentWithUsername, 0)

	err := q.db.Select(&comments, query, postId)
	if err != nil {
		return []CommentWithUsername{}, err
	}

	return comments, nil
}
