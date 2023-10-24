package test

import (
	"log"
	"os"
	"testing"

	"Github.com/Synoptic2023/internal/comment"
	"Github.com/Synoptic2023/internal/listing"
	"Github.com/Synoptic2023/internal/post"
	"Github.com/Synoptic2023/internal/user"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var UserRepo user.UserRepository
var PostRepo post.PostRepository
var ListRepo listing.ListingRepository
var CommentRepo comment.CommentRepository

func TestMain(m *testing.M) {
	db, err := sqlx.Connect("pgx", "postgresql://root:secret@localhost:5432/synoptictest?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	UserRepo = user.NewUserRepository(db)
	PostRepo = post.NewPosteRepository(db)
	ListRepo = listing.NewListingRepository(db)
	CommentRepo = comment.NewCommentRepository(db)

	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, user.Username, newUser.Username)
	require.Equal(t, "", newUser.Password)

	require.NotZero(t, newUser.ID)
	require.NotZero(t, newUser.CreatedAt)
}

func TestListUsers(t *testing.T) {
	users, err := UserRepo.ListUsers()
	require.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestGetUserByUsername(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	retrievedUser, err := UserRepo.GetUserByUsername(newUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, retrievedUser.Username, newUser.Username)
	require.Equal(t, retrievedUser.ID, newUser.ID)
	require.Equal(t, retrievedUser.CreatedAt, newUser.CreatedAt)
}

func TestGetUserById(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	retrievedUser, err := UserRepo.GetUserById(newUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, retrievedUser.Username, newUser.Username)
	require.Equal(t, retrievedUser.Password, newUser.Password)
	require.Equal(t, retrievedUser.ID, newUser.ID)
	require.Equal(t, retrievedUser.CreatedAt, newUser.CreatedAt)
}

func TestCreateListing(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	listing := NewListing(newUser.ID)
	newListing, err := ListRepo.CreateListing(listing)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, newListing.Username, newUser.Username)
	require.Equal(t, listing.Body, newListing.Body)

	require.NotZero(t, newListing.ID)
	require.NotZero(t, newListing.CreatedAt)
}

func TestCreatePost(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	post := NewPost(newUser.ID)
	newPost, err := PostRepo.CreatePost(post)
	require.NoError(t, err)
	require.NotEmpty(t, newPost)

	require.Equal(t, newUser.Username, newPost.Username)
	require.Equal(t, post.Body, newPost.Body)

	require.NotZero(t, newPost.ID)
	require.NotZero(t, newPost.CreatedAt)
}

func TestCreateComment(t *testing.T) {
	user := NewUser()
	newUser, err := UserRepo.CreateUser(user)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	post := NewPost(newUser.ID)
	newPost, err := PostRepo.CreatePost(post)
	require.NoError(t, err)
	require.NotEmpty(t, newPost)

	comment := NewComment(newUser.ID, newPost.ID)
	newComment, err := CommentRepo.CreateComment(comment)
	require.NoError(t, err)
	require.NotEmpty(t, newComment)

	require.Equal(t, newUser.Username, newComment.Username)
	require.Equal(t, comment.Body, newComment.Body)

	require.NotZero(t, newComment.ID)
	require.NotZero(t, newComment.CreatedAt)
}
