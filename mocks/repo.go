package mocks

import "gostarter/internal/data"

type repo struct {
}

func NewRepo() data.Repo {
	return &repo{}
}

func (r *repo) GetPosts() (*[]data.Post, error) {
	return &[]data.Post{{ID: "acb123", Caption: "Test"}}, nil
}

func (r *repo) GetPostByID(id string) (*data.Post, error) {
	return &data.Post{ID: "acb123", Caption: "Test"}, nil

}

func (r *repo) DeletePost(id string) error {
	return nil
}
