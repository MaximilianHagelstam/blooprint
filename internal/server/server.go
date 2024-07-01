package server

import (
	"gostarter/internal/database"
	"gostarter/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	App        *fiber.App
	Repository repository.Repository
	logger     log.AllLogger
}

func New() *Server {
	db := database.New()
	return &Server{
		App:        fiber.New(),
		Repository: repository.New(db),
		logger:     log.DefaultLogger(),
	}
}

func (s *Server) RegisterRoutes() {
	r := s.App.Group("/api")
	r.Get("/posts", s.GetPostsHandler)
	r.Post("/posts", s.CreatePostHandler)
}
