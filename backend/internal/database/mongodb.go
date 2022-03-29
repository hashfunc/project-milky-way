package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	client *mongo.Client
	config *Config
}

func (connection *Connection) DB() *mongo.Database {
	return connection.client.Database(connection.config.Name)
}

func (connection *Connection) Close() error {
	return connection.client.Disconnect(nil)
}

func New(config *Config) (*Connection, error) {
	client, err := mongo.Connect(nil, options.Client().ApplyURI(config.uri()))

	if err != nil {
		return nil, err
	}

	connection := &Connection{
		client: client,
		config: config,
	}

	return connection, nil
}
