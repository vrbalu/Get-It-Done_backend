package services

import (
	"GIT/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
)

const projectsCollection = "gid_projects"

func InsertProject() error {
	res, err := DB.InsertOne(projectsCollection, gin.H{"Hello": "World2"})
	if err != nil {
		return err
	}
	helpers.Log.Info(fmt.Sprintf("Insterted ID: %v", res.InsertedID))
	return nil
}
