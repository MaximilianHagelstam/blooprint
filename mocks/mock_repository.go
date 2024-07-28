package mocks

import (
	"blooprint/internal"
	"blooprint/internal/repository"
)

type mockRepository struct{}

func NewRepository() repository.Repository {
	return &mockRepository{}
}

func (r *mockRepository) GetPosts() (*[]internal.Post, error) {
	return &[]internal.Post{{ID: "acb123", Caption: "Test"}}, nil
}

func (r *mockRepository) CreatePost(post *internal.Post) error {
	return nil
}
