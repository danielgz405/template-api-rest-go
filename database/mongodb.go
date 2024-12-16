package database

// mongo db with atlas
import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	client *mongo.Client
}

func NewMongoRepo(url string) (*MongoRepo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return &MongoRepo{client: client}, nil
}

func (repo *MongoRepo) Close() error {
	return repo.client.Disconnect(context.Background())
}
