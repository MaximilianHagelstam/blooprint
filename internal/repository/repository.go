package repository

import (
	"gostarter/internal"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetPosts() (*[]internal.Post, error)
}

type repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) Repository {
	return &repository{db: db}
}
