package controllers

import (
	"GID/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type ProjectsController struct{}

const projectCollection = "projects"

func (*ProjectsController) GetAllProjects(c *gin.Context) {
	res, err := services.DB.Find(projectCollection, bson.M{}) // Verify functionality
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	if res == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (*ProjectsController) CreateProject(c *gin.Context) {
	err := services.InsertProject()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "OK")
}
func (*ProjectsController) UpdateProject(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (*ProjectsController) GetDescription(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
func (*ProjectsController) DeleteProject(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
