package controllers

import (
	"GID/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectsController struct{}

func (*ProjectsController) GetAllProjects(c *gin.Context) {

	c.JSON(http.StatusOK, "OK")
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
