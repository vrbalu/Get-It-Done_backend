package controllers

import (
	"GID/models"
	"GID/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProjectsController struct{}

func (*ProjectsController) GetAllProjects(c *gin.Context) {
	projects, err := services.GetAllProjects()
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	if projects == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, projects)
}
func (*ProjectsController) CreateProject(c *gin.Context) {
	var newProject models.Project
	if err := c.BindJSON(&newProject); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON structure for Project"})
		return
	}
	id, err := services.InsertProject(newProject)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
func (*ProjectsController) UpdateProject(c *gin.Context) {
	idOrKey := c.Param("projectIdOrKey")
	if idOrKey == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "idOrKey is empty"})
		return
	}
	project, err := services.GetOneProject(idOrKey)
	if err != nil {
		return
	}
	project.Id = ""
	if err := c.ShouldBindJSON(&project); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON structure for Project"})
		return
	}

	prjct, err := services.UpdateOneProject(idOrKey, project)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, prjct)
	return
}

func (*ProjectsController) GetDescription(c *gin.Context) {
	idOrKey := c.Param("projectIdOrKey")
	prjct, err := services.GetOneProject(idOrKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, prjct)
}
func (*ProjectsController) DeleteProject(c *gin.Context) {
	idOrKey := c.Param("projectIdOrKey")
	err := services.DeleteProject(idOrKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}
