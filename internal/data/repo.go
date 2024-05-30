package data

import (
	"database/sql"
)

type Repo interface {
	GetPosts() (*[]Post, error)
	GetPostByID(id string) (*Post, error)
	DeletePost(id string) error
}

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetPosts() (*[]Post, error) {
	rows, err := r.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.Caption); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func (r *repo) GetPostByID(id string) (*Post, error) {
	row := r.db.QueryRow("SELECT id, caption FROM posts WHERE id = $1", id)
	post := Post{}
	err := row.Scan(&post.ID, &post.Caption)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repo) DeletePost(id string) error {
	_, err := r.db.Query("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
