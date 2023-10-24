package post

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostRepository interface {
	CreatePost(post Post) (PostWithUsername, error)
	GetPost(postId int64) (PostWithUsername, error)
	ListPosts() ([]PostWithUsername, error)
}

type postQueries struct {
	db *sqlx.DB
}

func NewPosteRepository(db *sqlx.DB) PostRepository {
	return &postQueries{db: db}
}

func (q *postQueries) CreatePost(post Post) (PostWithUsername, error) {
	query := "INSERT INTO posts (title, body, user_id) VALUES ($1, $2, $3) RETURNING *"
	newPost := PostWithUsername{}

	err := q.db.Get(&newPost, query, post.Title, post.Body, post.UserID)
	if err != nil {
		return PostWithUsername{}, fmt.Errorf("could not create new post %w", err)
	}

	fetchUsername := "SELECT posts.*, users.username FROM posts JOIN users ON posts.user_id = users.id WHERE posts.id = $1"
	err = q.db.Get(&newPost, fetchUsername, newPost.ID)
	if err != nil {
		return PostWithUsername{}, fmt.Errorf("could not retrieve new post %w", err)
	}

	return newPost, nil
}

func (q *postQueries) GetPost(postId int64) (PostWithUsername, error) {
	query := "SELECT posts.*, users.username FROM posts JOIN users ON posts.user_id = users.id WHERE posts.id = $1"
	post := PostWithUsername{}
	err := q.db.Get(&post, query, postId)
	if err != nil {
		return PostWithUsername{}, err
	}

	return post, nil
}

func (q *postQueries) ListPosts() ([]PostWithUsername, error) {
	query := "SELECT posts.*, users.username FROM posts JOIN users ON posts.user_id = users.id"
	posts := make([]PostWithUsername, 0)

	err := q.db.Select(&posts, query)
	if err != nil {
		return []PostWithUsername{}, fmt.Errorf("could not find any posts %w", err)
	}

	return posts, nil
}
