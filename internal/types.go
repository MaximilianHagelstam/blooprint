package internal

type Post struct {
	ID      string `bson:"id"`
	Caption string `bson:"caption,omitempty"`
}
