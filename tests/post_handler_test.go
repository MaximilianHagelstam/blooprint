package tests

import (
	"blooprint/internal/server"
	"blooprint/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTestServer() *server.Server {
	return &server.Server{
		Repository: mocks.NewRepository(),
	}
}

func TestGetPostsHandler(t *testing.T) {
	s := SetupTestServer()
	server := httptest.NewServer(http.HandlerFunc(s.GetPostsHandler))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", resp.Status)
	}
}

func TestCreatePostHandler(t *testing.T) {
	s := SetupTestServer()
	server := httptest.NewServer(http.HandlerFunc(s.CreatePostHandler))
	defer server.Close()

	postData := map[string]string{"caption": "Test Post"}
	jsonData, _ := json.Marshal(postData)

	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %v", resp.Status)
	}
}
