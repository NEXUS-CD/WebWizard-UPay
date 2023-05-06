package database

import (
	"context"
	"log"
	"sync"
	"time"

	"UPay/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *mongo.Database
var mongoCtx context.Context
var mongoOnce sync.Once

func InitMongoDB() {
	mongoOnce.Do(func() {
		conf := configs.GetConfig().MongoDB
		uri := "mongodb://" + conf.Host + ":" + conf.Port
		clientOptions := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		mongoInstance = client.Database(conf.Database)
		mongoCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	})
}

// GetCollection 返回给定名称的集合
func GetCollection(collectionName string) *mongo.Collection {
	return mongoInstance.Collection(collectionName)
}

func GetMongoCtx() context.Context {
	return mongoCtx
}
