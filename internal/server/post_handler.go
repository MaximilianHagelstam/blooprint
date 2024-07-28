package server

import (
	"blooprint/internal"
	"encoding/json"
	"net/http"
)

func (s *Server) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.Repository.GetPosts()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(posts)
}

func (s *Server) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post internal.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil || post.Caption == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Repository.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
