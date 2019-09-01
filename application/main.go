package application

import (
	"context"
	"messenger/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error
var ctx context.Context
var dbName string

func init() {
	conf, err := config.Get("mongodb")

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(conf["connection"].(string)))
	dbName = conf["dbName"].(string)

	if err != nil {
		panic(err)
	}
}
