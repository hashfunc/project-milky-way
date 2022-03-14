package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Context struct {
	client   *mongo.Client
	database *mongo.Database
}

func (connection *Context) DB(name string, opts ...*options.DatabaseOptions) (*mongo.Database, error) {
	if connection.client == nil {
		return nil, errors.New("mongodb 연결 실패")
	}

	db := connection.client.Database(name, opts...)

	if db == nil {
		return nil, errors.New("database 조회 실패")
	}

	return db, nil
}

func (connection *Context) Close() error {
	return connection.client.Disconnect(nil)
}

var (
	Connection *Context
)

func Connect(config *Config) error {
	uri := config.uri()

	client, err := mongo.Connect(nil, options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	Connection = &Context{
		client: client,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return nil
}
