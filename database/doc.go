package database

import (
	"context"
	"example/web-service-gin/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Database = nil

// GetConnection is for get mongo connection
func GetConnection() *mongo.Database {
	if collection == nil {
		ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		connectString := utils.EnvVar("DB_CONNECTION_STRING", "")
		dbName := utils.EnvVar("DB_NAME", "")

		clientOptions := options.Client().ApplyURI(connectString)
		client, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			log.Fatal(err)
		} else {
			collection = client.Database(dbName)
			return collection
		}
	}
	return collection
}
