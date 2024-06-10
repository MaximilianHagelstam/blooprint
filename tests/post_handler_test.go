package handlers

import (
	"gostarter/internal/handlers"
	"gostarter/mocks"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetPostsHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/posts", handlers.GetPostsHandler(&handlers.PostHandlerConfig{
		PostRepository: mocks.NewPostRepository(),
	}))

	req, err := http.NewRequest("GET", "/api/v1/posts", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200, got %v", resp.Status)
	}
}

func TestGetPostByIDHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/posts/:id", handlers.GetPostByIDHandler(&handlers.PostHandlerConfig{
		PostRepository: mocks.NewPostRepository(),
	}))

	req, err := http.NewRequest("GET", "/api/v1/posts/123", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200, got %v", resp.Status)
	}
}

func TestDeletePostHandler(t *testing.T) {
	app := fiber.New()
	app.Delete("/api/v1/posts/:id", handlers.DeletePostHandler(&handlers.PostHandlerConfig{
		PostRepository: mocks.NewPostRepository(),
	}))

	req, err := http.NewRequest("DELETE", "/api/v1/posts/123", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server: %v", err)
	}

	if resp.StatusCode != fiber.StatusNoContent {
		t.Errorf("expected status 204, got %v", resp.Status)
	}
}
