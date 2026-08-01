package database

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

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

func InsertNotification(ctx context.Context, mClient *mongo.Client, notification entity.Notification) error {
	coll := mClient.Database(env.GetString("DATABASE_NAME", "OrderyNotifications")).Collection(env.GetString("COLLECTION_NAME", "Notifications"))

	_, err := coll.InsertOne(ctx, notification)

	return err
}

func GetNotificationsPagination(ctx context.Context, mClient *mongo.Client, userId string, page, limit int64) (*entity.NotificationResult, error) {

	coll := mClient.Database(env.GetString("DATABASE_NAME", "OrderyNotifications")).Collection(env.GetString("COLLECTION_NAME", "Notifications"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	queryFilter := bson.M{"userid": userId}

	totalCount, err := coll.CountDocuments(ctx, queryFilter)

	if err != nil {
		return nil, err
	}

	sortFilter := bson.D{{Key: "createdat", Value: -1}}

	ops := options.Find().SetSort(sortFilter).SetSkip(offset).SetLimit(limit)

	cursur, err := coll.Find(context.Background(), queryFilter, ops)

	if err != nil {
		return nil, err
	}

	defer cursur.Close(ctx)

	var notification []entity.Notification

	if err := cursur.All(ctx, &notification); err != nil {
		return nil, err
	}

	nResult := entity.NotificationResult{
		Notifications: notification,
		Total:         totalCount,
		Page:          page,
		Limit:         limit,
	}

	return &nResult, nil
}

func GetNotification(ctx context.Context, id string, mClient *mongo.Client) (*entity.Notification, error) {
	println("Id sent: " + id)
	coll := mClient.Database(env.GetString("DATABASE_NAME", "OrderyNotifications")).Collection(env.GetString("COLLECTION_NAME", "Notifications"))

	filter := bson.M{"id": id}

	cursor := coll.FindOne(ctx, filter)

	var notification entity.Notification

	// if err != nil {
	// 	return nil, err
	// }

	// defer cursor.Close(ctx)
	r, _ := cursor.Raw()
	raw := r.String()
	data := bytes.NewBufferString(raw)
	fmt.Println("bson raw:: ", r.String())

	err := json.Unmarshal(data.Bytes(), &notification)
	if err != nil {
		return nil,err
	}

	// fmt.Println("bson raw:: ", r.String())
	// if err := cursor.Decode(notification); err != nil {
	// 	return nil, err
	// }

	return &notification, nil

}
