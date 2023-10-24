package listing

type ListingService interface {
	CreateListing(title string, body string, userId int64) (ListingWithUsername, error)
	GetListing(listingId int64) (ListingWithUsername, error)
	CloseListing(listingId int64) (ListingWithUsername, error)
	ListListings() ([]ListingWithUsername, error)
}

type listingService struct {
	repo ListingRepository
}

func NewListingService(repo ListingRepository) ListingService {
	return &listingService{repo: repo}
}

func (s *listingService) CreateListing(title string, body string, userId int64) (ListingWithUsername, error) {

	listing := ListingWithUsername{Title: title, Body: body, UserID: userId}
	newListing, err := s.repo.CreateListing(listing)
	if err != nil {
		return ListingWithUsername{}, err
	}

	return newListing, nil

}

func (s *listingService) GetListing(listingId int64) (ListingWithUsername, error) {
	listing, err := s.repo.GetListing(listingId)
	if err != nil {
		return ListingWithUsername{}, err
	}

	return listing, nil

}

func (s *listingService) ListListings() ([]ListingWithUsername, error) {
	listings, err := s.repo.ListListings()
	if err != nil {
		return []ListingWithUsername{}, err
	}

	return listings, nil
}

func (s *listingService) CloseListing(listingId int64) (ListingWithUsername, error) {
	listing, err := s.repo.CloseListing(listingId)
	if err != nil {
		return ListingWithUsername{}, err
	}

	return listing, nil
}
