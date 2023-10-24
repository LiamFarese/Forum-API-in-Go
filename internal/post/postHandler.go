package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	ListPosts(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	service PostService
}

func NewPostHandler(service PostService) PostHandler {
	return &postHandler{service: service}
}

func (h *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var params CreatePostParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := h.service.CreatePost(params.Title, params.Body, params.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}

func (h *postHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	postId, err := strconv.ParseInt(chi.URLParam(r, "listingId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPost(postId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)
}

func (h *postHandler) ListPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := h.service.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(posts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}
