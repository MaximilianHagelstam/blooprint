package tests

import (
	"gostarter/internal/handlers"
	"gostarter/mocks"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetPostsHandler(t *testing.T) {
	app := fiber.New()
	repo := mocks.NewRepo()
	app.Get("/api/posts", handlers.GetPostsHandler(repo))

	req, err := http.NewRequest("GET", "/api/posts", nil)
	if err != nil {
		t.Fatalf("error creating request. Err: %v", err)
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
}
