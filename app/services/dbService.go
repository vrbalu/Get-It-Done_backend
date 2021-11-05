package services

import (
	"GID/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (ds *DbServiceType) establishConnection() {
	clientOptions := options.Client().
		ApplyURI(config.DBUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ds.db = client.Database(config.DBName)
	return
}

func (ds DbServiceType) InsertOne(coll string, data interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	one, err := ds.db.Collection(coll).InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (ds DbServiceType) GetOne(coll string, key string, value string, resultModel interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{key, value}}
	err := ds.db.Collection(coll).FindOne(ctx, filter).Decode(resultModel)
	if err != nil {
		return nil, err
	}
	return resultModel, nil
}

func (ds DbServiceType) FindByField(coll string, field string, value string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter, err := buildFilter(field, value)
	if err != nil {
		defer cancel()
		return nil, err
	}
	var result []bson.M
	cursor, err := ds.db.Collection(coll).Find(ctx, filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
	}(cursor, ctx)
	if err != nil {
		defer cancel()
		return nil, err
	}

	for cursor.Next(ctx) {
		var item bson.M
		err = cursor.Decode(&item)
		if err == nil {
			result = append(result, item)
		}
	}
	defer cancel()
	return result, nil
}

func buildFilter(field string, value string) (bson.M, error) {
	if field == "_id" {
		objId, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			return nil, err
		}
		return bson.M{field: objId}, nil
	}
	return bson.M{field: value}, nil
}

func (ds DbServiceType) DeleteOne(coll string, key string, value string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{key, value}}
	res, err := ds.db.Collection(coll).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ds DbServiceType) UpdateOneById(coll string, id string, updateValue interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := ds.db.Collection(coll).UpdateByID(ctx, id, updateValue)
	if err != nil {
		return nil, err
	}
	return res, nil
}
