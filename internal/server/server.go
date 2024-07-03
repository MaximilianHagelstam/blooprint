package server

import (
	"fmt"
	"gostarter/internal/database"
	"gostarter/internal/repository"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	Repository repository.Repository
	port       int
	logger     *log.Logger
}

func New() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	logger := log.Default()
	db := database.New()
	repo := repository.New(db)

	NewServer := &Server{
		Repository: repo,
		port:       port,
		logger:     logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/posts", s.GetPostsHandler)
	r.Post("/api/posts", s.CreatePostHandler)

	return r
}
