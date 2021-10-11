package services

import (
	"GIT/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (db DbServiceType) InsertOne(coll string, data interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	one, err := db.db.Collection(coll).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (db DbServiceType) GetOne(coll string, key string, value string, resultModel interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{key, value}}
	err := db.db.Collection(coll).FindOne(ctx, filter).Decode(resultModel)
	if err != nil {
		return nil, err
	}
	return resultModel, nil
}

func (db DbServiceType) DeleteOne(coll string, key string, value string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{key, value}}
	res, err := db.db.Collection(coll).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db DbServiceType) UpdateOneById(coll string, id string, updateValue interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := db.db.Collection(coll).UpdateByID(ctx, id, updateValue)
	if err != nil {
		return nil, err
	}
	return res, nil
}
