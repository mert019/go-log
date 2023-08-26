package logdestinations

import (
	"context"

	"github.com/mert019/go-log/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBLoggerConfiguration struct {
	DbUri      string
	Database   string
	Collection string
}

func NewMongoDBLogger(configuration MongoDBLoggerConfiguration) (core.ILogDestination, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configuration.DbUri))
	if err != nil {
		return nil, &core.LogDestinationConnectionError{
			Destination:     "MongoDB",
			ConnectionError: err,
		}
	}
	collection := client.Database(configuration.Database).Collection(configuration.Collection)

	return &MongoDBLogger{
		configuration: configuration,
		client:        client,
		collection:    collection,
	}, nil
}

type MongoDBLogger struct {
	configuration MongoDBLoggerConfiguration
	client        *mongo.Client
	collection    *mongo.Collection
}

func (mongoDBLogger *MongoDBLogger) Log(log core.Log) error {
	_, err := mongoDBLogger.collection.InsertOne(context.TODO(), log)
	return err
}

func (mongoDBLogger *MongoDBLogger) Close() error {
	return mongoDBLogger.client.Disconnect(context.TODO())
}
