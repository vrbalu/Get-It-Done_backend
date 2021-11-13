package services

import (
	"GID/helpers"
	"GID/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	projectsCollection = "projects"
)

var dbKeyOrId = [2]string{"name", "_id"}

func GetOneProject(idOrKey string) (*models.Project, error) {
	if idOrKey == "" {
		return nil, errors.New("Key or ID is empty|")
	}
	for _, s := range dbKeyOrId {
		r, err := DB.GetOne(projectsCollection, s, idOrKey)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return nil, err
			}
		}
		if r != nil {
			project, err := marshalToProject(r)
			if err != nil {
				return nil, err
			}
			return project, nil
		}
	}
	return nil, errors.New("no results found")
}

func UpdateOneProject(idOrKey string, data interface{}) (interface{}, error) {
	if idOrKey == "" {
		return nil, errors.New("Key or ID is empty|")
	}
	for _, s := range dbKeyOrId {
		r, err := DB.UpdateOne(projectsCollection, s, idOrKey, data)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return nil, err
			}
		}
		if r.ModifiedCount > 0 {
			return data, nil
		}
	}
	return nil, errors.New("no results found")
}

func GetAllProjects() ([]models.Project, error) {
	var prjcts []models.Project
	res, err := DB.Find(projectsCollection, bson.M{})
	if err != nil {
		return nil, err
	}
	for _, m := range res {
		project, err := marshalToProject(&m)
		if err != nil {
			return nil, err
		}
		prjcts = append(prjcts, *project)
	}
	return prjcts, nil
}

func InsertProject(project models.Project) (interface{}, error) {
	res, err := DB.InsertOne(projectsCollection, project)
	if err != nil {
		return nil, err
	}
	helpers.Log.Info(fmt.Sprintf("Insterted ID: %v", res.InsertedID))
	return res.InsertedID, nil
}

func DeleteProject(idOrKey string) error {
	for _, s := range dbKeyOrId {
		r, err := DB.DeleteOne(projectsCollection, s, idOrKey)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return err
			}
		}
		if r.DeletedCount > 0 {
			return nil
		}
	}
	return errors.New("not deleted")
}

func marshalToProject(m *bson.M) (*models.Project, error) {
	var proj *models.Project

	b, err := bson.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(b, &proj)
	if err != nil {
		return nil, err
	}
	return proj, nil
}
