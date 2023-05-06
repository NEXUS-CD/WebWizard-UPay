package models

import (
	"UPay/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection 返回给定名称的集合
func coll(name string) *mongo.Collection {
	return database.GetCollection(name)
}
