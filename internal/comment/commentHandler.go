package comment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CommentHandler interface {
	CreateComment(w http.ResponseWriter, r *http.Request)
	Reply(w http.ResponseWriter, r *http.Request)
	GetCommentsFromPost(w http.ResponseWriter, r *http.Request)
}

type commentHandler struct {
	service CommentService
}

func NewCommentHandler(service CommentService) CommentHandler {
	return &commentHandler{service: service}
}

func (h *commentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var params CreateCommentParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := h.service.CreateComment(params.Body, params.UserID, params.PostID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}

func (h *commentHandler) Reply(w http.ResponseWriter, r *http.Request) {
	var params CreateReplyParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := h.service.Reply(params.Body, params.UserID, params.PostID, params.ParentComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRep, err := json.Marshal(comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)

}

func (h *commentHandler) GetCommentsFromPost(w http.ResponseWriter, r *http.Request) {

	postId, err := strconv.ParseInt(chi.URLParam(r, "listingId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := h.service.GetCommentsFromPost(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonRep, err := json.Marshal(comments)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRep)
}
