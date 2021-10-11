package services

import (
	"GIT/helpers"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InsertProject() error {
	ctx := context.TODO()
	mongoService := DB.db.Collection("testing")
	res, err := mongoService.InsertOne(ctx, gin.H{"Hello": "World"})
	if err != nil {
		return err
	}
	helpers.Log.Info(fmt.Sprintf("Insterted ID: %v", res.InsertedID))
	return nil
}
