package listing

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ListingHandler interface {
	CreateListing(w http.ResponseWriter, r *http.Request)
	GetListing(w http.ResponseWriter, r *http.Request)
	CloseListing(w http.ResponseWriter, r *http.Request)
	ListListings(w http.ResponseWriter, r *http.Request)
}

type listingHandler struct {
	service ListingService
}

func NewListingHandler(service ListingService) ListingHandler {
	return &listingHandler{service: service}
}

func (h *listingHandler) CreateListing(w http.ResponseWriter, r *http.Request) {
	var params CreateListingParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	listing, err := h.service.CreateListing(params.Title, params.Body, params.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(listing)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}

func (h *listingHandler) GetListing(w http.ResponseWriter, r *http.Request) {

	listingId, err := strconv.ParseInt(chi.URLParam(r, "listingId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listing, err := h.service.GetListing(listingId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(listing)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)
}

func (h *listingHandler) CloseListing(w http.ResponseWriter, r *http.Request) {

	listingId, err := strconv.ParseInt(chi.URLParam(r, "listingId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listing, err := h.service.CloseListing(listingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)
}

func (h *listingHandler) ListListings(w http.ResponseWriter, r *http.Request) {

	listings, err := h.service.ListListings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(listings)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}
