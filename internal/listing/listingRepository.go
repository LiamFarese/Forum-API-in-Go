package listing

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ListingRepository interface {
	CreateListing(listing ListingWithUsername) (ListingWithUsername, error)
	GetListing(listingId int64) (ListingWithUsername, error)
	CloseListing(listingId int64) (ListingWithUsername, error)
	ListListings() ([]ListingWithUsername, error)
}

type listingQueries struct {
	db *sqlx.DB
}

func NewListingRepository(db *sqlx.DB) ListingRepository {
	return &listingQueries{db: db}
}

func (q *listingQueries) CreateListing(listing ListingWithUsername) (ListingWithUsername, error) {

	query := "INSERT INTO listings (title, body, user_id) VALUES ($1, $2, $3) RETURNING *"
	newListing := ListingWithUsername{}

	err := q.db.Get(&newListing, query, listing.Title, listing.Body, listing.UserID)
	if err != nil {
		return ListingWithUsername{}, fmt.Errorf("could not create new listing %w", err)
	}

	fetchUsername := "SELECT listings.*, users.username FROM listings JOIN users ON listings.user_id = users.id WHERE listings.id = $1"
	err = q.db.Get(&newListing, fetchUsername, newListing.ID)
	if err != nil {
		return ListingWithUsername{}, fmt.Errorf("could not retrieve new listing %w", err)
	}

	return newListing, nil
}

func (q *listingQueries) GetListing(listingId int64) (ListingWithUsername, error) {
	query := "SELECT listings.*, users.username FROM listings JOIN users ON listings.user_id = users.id WHERE listings.id = $1"
	listing := ListingWithUsername{}
	err := q.db.Get(&listing, query, listingId)
	if err != nil {
		return ListingWithUsername{}, err
	}

	return listing, nil
}

func (q *listingQueries) ListListings() ([]ListingWithUsername, error) {
	query := "SELECT listings.*, users.username FROM listings JOIN users ON listings.user_id = users.id WHERE listings.is_settled = FALSE"
	listings := make([]ListingWithUsername, 0)

	err := q.db.Select(&listings, query)
	if err != nil {
		return []ListingWithUsername{}, fmt.Errorf("could not find any listings %w", err)
	}

	return listings, nil
}

func (q *listingQueries) CloseListing(listingId int64) (ListingWithUsername, error) {
	query := "UPDATE listings SET is_settled = TRUE WHERE id = $1"

	_, err := q.db.Exec(query, listingId)
	if err != nil {
		return ListingWithUsername{}, err
	}

	listing, err := q.GetListing(listingId)
	if err != nil {
		return ListingWithUsername{}, err
	}

	return listing, nil
}
