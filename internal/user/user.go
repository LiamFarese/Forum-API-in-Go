package user

import (
	"Github.com/Synoptic2023/internal/listing"
	"Github.com/Synoptic2023/internal/post"
	"github.com/jackc/pgx/pgtype"
)

type User struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	// hashed password
	Password  string             `db:"password" json:"password"`
	Role      string             `db:"role" json:"role"`
	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

type createUserInput struct {
	Username string
	Password string
	Role     string
}

type logininput struct {
	Username string
	Password string
}

type Profile struct {
	User     User
	Listings []listing.ListingWithUsername
	Posts    []post.PostWithUsername
}
