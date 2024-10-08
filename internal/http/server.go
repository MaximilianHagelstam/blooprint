package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/maximilianhagelstam/blooprint/internal/database"
	v1 "github.com/maximilianhagelstam/blooprint/internal/v1"
)

type Server struct {
	handler *v1.Handler
	port    int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db := database.New()
	repository := v1.NewRepository(db)
	handler := v1.NewHandler(repository)

	server := &Server{
		port:    port,
		handler: handler,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", server.port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/v1/posts", s.handler.GetPostsHandler)
	r.Post("/api/v1/posts", s.handler.CreatePostHandler)
	r.Delete("/api/v1/posts/{id}", s.handler.DeletePostHandler)

	return r
}
