package test

import (
	"math/rand"
	"strings"
	"time"

	"Github.com/Synoptic2023/internal/comment"
	"Github.com/Synoptic2023/internal/listing"
	"Github.com/Synoptic2023/internal/post"
	"Github.com/Synoptic2023/internal/user"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func NewUser() user.User {
	username := RandomString(6)
	password := RandomString(8)
	role := RandomString(5)

	user := user.User{Username: username, Password: password, Role: role}
	return user
}

func NewListing(userId int64) listing.ListingWithUsername {

	title := RandomString(10)
	body := RandomString(140)

	listing := listing.ListingWithUsername{Title: title, Body: body, UserID: userId}
	return listing
}

func NewPost(userId int64) post.Post {
	title := RandomString(10)
	body := RandomString(1000)

	post := post.Post{Title: title, Body: body, UserID: userId}
	return post
}

func NewComment(userId int64, postId int64) comment.Comment {
	body := RandomString(200)
	return comment.Comment{Body: body, UserID: userId, PostID: postId}

}
