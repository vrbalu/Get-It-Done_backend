package services

import (
	"GIT/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DB DbServiceType

type DbServiceType struct {
	db *mongo.Database
}

func init() {
	DB.establishConnection()
}
func (db *DbServiceType) establishConnection() {
	clientOptions := options.Client().
		ApplyURI(config.DBUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db.db = client.Database(config.DBName)
	return
}
