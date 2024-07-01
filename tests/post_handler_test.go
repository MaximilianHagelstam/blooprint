package handlers

import (
	"gostarter/internal/server"
	"gostarter/mocks"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func SetupTestServer() *server.Server {
	return &server.Server{
		App:        fiber.New(),
		Repository: mocks.NewRepository(),
	}
}

func TestGetPostsHandler(t *testing.T) {
	s := SetupTestServer()
	s.App.Get("/api/posts", s.GetPostsHandler)

	req, err := http.NewRequest("GET", "/api/posts", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	resp, err := s.App.Test(req)
	if err != nil {
		t.Fatalf("error making request to server: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200, got %v", resp.Status)
	}
}
