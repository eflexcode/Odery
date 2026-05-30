package database

import (
	"context"

	"github.com/cmd/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DatabaseRep struct{
	Mongo *mongo.Client
}

func ConnDatabase(config config.DatabaseConfig,ctx context.Context)(*mongo.Client,error){
	var client,_ = mongo.Connect(options.Client().ApplyURI(config.ConnUrl))
	err := client.Ping(ctx,readpref.Primary())
	return client,err
}
