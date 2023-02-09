package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	client *mongo.Client
}

func (client *Client) Connect(ctx context.Context) error {
	return client.client.Connect(ctx)
}

func (client *Client) Disconnect(ctx context.Context) error {
	return client.client.Disconnect(ctx)
}

func (client *Client) GetCollection(name, collection string) *mongo.Collection {
	return client.client.Database(name).Collection(collection)
}

func (client *Client) Ping() error {
	return client.client.Ping(context.Background(), readpref.Primary())
}

func NewClient(uri string) (*Client, error) {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(opts)

	if err != nil {
		return nil, err
	}

	instance := &Client{client: client}

	return instance, nil
}
