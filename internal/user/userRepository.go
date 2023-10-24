package user

import (
	"fmt"

	"Github.com/Synoptic2023/internal/listing"
	"Github.com/Synoptic2023/internal/post"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserById(userId int64) (User, error)
	GetUserProfile(user User) (Profile, error)
	ListUsers() ([]User, error)
}

type userQueries struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userQueries{db: db}
}

func (q *userQueries) CreateUser(user User) (User, error) {
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id, username, role, created_at"
	newUser := User{}
	err := q.db.Get(&newUser, query, user.Username, user.Password, user.Role)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (q *userQueries) GetUserByUsername(username string) (User, error) {
	query := "SELECT * FROM users WHERE username = $1 LIMIT 1;"
	user := User{}

	err := q.db.Get(&user, query, username)
	if err != nil {
		return User{}, fmt.Errorf("User does not exist %w", err)
	}

	return user, nil
}

func (q *userQueries) GetUserById(userId int64) (User, error) {
	query := "SELECT * FROM users WHERE id = $1 LIMIT 1;"
	user := User{}

	err := q.db.Get(&user, query, userId)
	if err != nil {
		return User{}, fmt.Errorf("User does not exist %w", err)
	}

	user.Password = ""

	return user, nil
}

func (q *userQueries) GetUserProfile(queryUser User) (Profile, error) {

	user, err := q.GetUserById(queryUser.ID)
	if err != nil {
		return Profile{}, err
	}

	listingQuery := "SELECT listings.*, users.username FROM listings JOIN users ON listings.user_id = users.id WHERE listings.user_id = $1"
	listings := make([]listing.ListingWithUsername, 0)
	err = q.db.Select(&listings, listingQuery, user.ID)
	if err != nil {
		return Profile{}, fmt.Errorf("could not query listings %w", err)
	}

	postQuery := "SELECT posts.*, users.username FROM posts JOIN users ON posts.user_id = users.id WHERE posts.user_id = $1"
	posts := make([]post.PostWithUsername, 0)
	err = q.db.Select(&posts, postQuery, user.ID)
	if err != nil {
		return Profile{}, fmt.Errorf("could not query posts %w", err)
	}

	profile := Profile{user, listings, posts}

	return profile, nil
}

func (q *userQueries) ListUsers() ([]User, error) {
	query := "SELECT id, username, role, created_at FROM users;"
	users := make([]User, 0)

	err := q.db.Select(&users, query)
	if err != nil {
		return []User{}, err
	}

	return users, nil

}
