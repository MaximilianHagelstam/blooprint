package database

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = os.Getenv("DB_DATABASE")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
)

func New() *mongo.Database {
	credential := options.Credential{
		Username: username,
		Password: password,
	}
	connStr := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOpts := options.Client().ApplyURI(connStr).SetAuth(credential)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to db: %s", err))
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(fmt.Sprintf("cannot ping db: %s", err))
	}

	return client.Database(database)
}
