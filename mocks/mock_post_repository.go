package mocks

import (
	"gostarter/internal"
	"gostarter/internal/repository"
)

type postRepository struct{}

func NewPostRepository() repository.PostRepository {
	return &postRepository{}
}

func (p *postRepository) GetPosts() (*[]internal.Post, error) {
	return &[]internal.Post{{ID: "acb123", Caption: "Test"}}, nil
}

func (p *postRepository) GetPostByID(id string) (*internal.Post, error) {
	return &internal.Post{ID: "acb123", Caption: "Test"}, nil
}

func (p *postRepository) DeletePost(id string) error {
	return nil
}
