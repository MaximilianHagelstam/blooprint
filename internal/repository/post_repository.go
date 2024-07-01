package repository

import (
	"context"
	"gostarter/internal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) GetPosts() (*[]internal.Post, error) {
	opts := options.Find().SetSort(bson.D{{Key: "date_ordered", Value: 1}})
	cursor, err := r.db.Collection("posts").Find(context.Background(), bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}

	posts := []internal.Post{}
	if err = cursor.All(context.Background(), &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *repository) CreatePost(post *internal.Post) error {
	_, err := r.db.Collection("posts").InsertOne(context.Background(), *post)
	if err != nil {
		return err
	}
	return nil
}
