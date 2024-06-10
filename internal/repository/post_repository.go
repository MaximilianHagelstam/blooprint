package repository

import (
	"database/sql"
	"gostarter/internal"
)

type PostRepository interface {
	GetPosts() (*[]internal.Post, error)
	GetPostByID(id string) (*internal.Post, error)
	DeletePost(id string) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) GetPosts() (*[]internal.Post, error) {
	rows, err := p.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []internal.Post{}
	for rows.Next() {
		post := internal.Post{}
		if err := rows.Scan(&post.ID, &post.Caption); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func (p *postRepository) GetPostByID(id string) (*internal.Post, error) {
	row := p.db.QueryRow("SELECT id, caption FROM posts WHERE id = $1", id)
	post := internal.Post{}
	err := row.Scan(&post.ID, &post.Caption)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *postRepository) DeletePost(id string) error {
	_, err := p.db.Query("DELETE FROM posts WHERE id = $1", id)
	return err
}
