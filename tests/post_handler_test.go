package handlers

import (
	"gostarter/internal/handlers"
	"gostarter/mocks"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func SetupTests() *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/posts", handlers.GetPostsHandler(&handlers.PostHandlerConfig{
		PostRepository: mocks.NewPostRepository(),
	}))
	return app
}

func TestGetPostsHandler(t *testing.T) {
	app := SetupTests()

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
