package db

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

func (connection *Context) setDatabase(
	name string, opts ...*options.DatabaseOptions,
) error {
	if connection.client == nil {
		return errors.New("db.Context: 클라이언트가 설정되지 않았습니다")
	}

	database := connection.client.Database(name, opts...)
	connection.database = database

	return nil
}

func (connection *Context) DB() *mongo.Database {
	return connection.database
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

	if err := Connection.setDatabase(config.Name); err != nil {
		panic(err)
	}

	return nil
}
