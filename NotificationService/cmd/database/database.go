package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repo struct {
	databaseMongo *mongo.Client
}

func ConnectDatabase(ctx context.Context, connUrl string) (*mongo.Client, error) {

	var cli, _ = mongo.Connect(options.Client().ApplyURI(connUrl))
	err := cli.Ping(ctx, readpref.Primary())

	return cli, err
}
