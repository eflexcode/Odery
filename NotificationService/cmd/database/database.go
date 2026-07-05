package database

import (
	"context"

	"github.com/cmd/entity"
	"github.com/env"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Repo struct {
	DatabaseMongo *mongo.Client
}

func ConnectDatabase(ctx context.Context, connUrl string) (*mongo.Client, error) {

	var cli, _ = mongo.Connect(options.Client().ApplyURI(connUrl))
	err := cli.Ping(ctx, readpref.Primary())

	return cli, err
}

func InsertNotification(){
	
}

func GetNotificationsPagination(ctx context.Context, mClient *mongo.Client, userId string, page, limit int64) ([]entity.Notification, error) {

	coll := mClient.Database(env.GetString("DATABASE_NAME", "OrderyNotifications")).Collection(env.GetString("COLLECTION_NAME", "Notifications"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	queryFilter := bson.M{"userId": userId}

	totalCount, err := coll.CountDocuments(ctx, queryFilter)

	if err != nil {
		return nil, err
	}

	sortFilter := bson.D{{key: "created_at", Value: -1}}

	ops := options.Find().SetSort(sortFilter).SetSkip(offset).SetLimit(limit)

	cursur, err := coll.Find(context.Background(), queryFilter, ops)
	
	if err  != nil{
		return nil,err
	}
	
	defer cursur.Close(ctx)
	
	var notification []entity.Notification 
	
	
	
}
