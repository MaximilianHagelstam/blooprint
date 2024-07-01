package mocks

import (
	"gostarter/internal"
	"gostarter/internal/repository"
)

type mockRepository struct{}

func NewRepository() repository.Repository {
	return &mockRepository{}
}

func (r *mockRepository) GetPosts() (*[]internal.Post, error) {
	return &[]internal.Post{{ID: "acb123", Caption: "Test"}}, nil
}
