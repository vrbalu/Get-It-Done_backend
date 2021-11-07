package controllers

import (
	"GID/models"
	"GID/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
)

// NOTE: In other collections will be reference only by ID
type TeamsController struct{}

const teamsCollection = "teams"

func (tc TeamsController) CreateTeam(c *gin.Context) {
	var team, existing *models.Team
	b, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(b, &team)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	err = services.DB.GetOne(teamsCollection, "name", team.Name, &existing)
	if err != mongo.ErrNoDocuments {
		c.AbortWithStatusJSON(500, err)
		return
	}
	if existing != nil {
		c.AbortWithStatusJSON(409, "A team with this name already exists, please pick another name.")
		return
	}
	one, err := services.DB.InsertOne(teamsCollection, &team)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": one.InsertedID})
}

func (tc TeamsController) GetMembers(c *gin.Context) {
	var team models.Team
	teamName := c.Param("team")
	err := services.DB.GetOne(teamsCollection, "name", teamName, &team)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"members": team.Members})

}
func (tc TeamsController) UpdateTeam(c *gin.Context) {
	var team models.Team
	teamName := c.Param("team")
	b, err := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(b, &team)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	res, err := services.DB.UpdateOne(teamsCollection, "name", teamName, team)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func (tc TeamsController) DeleteTeam(c *gin.Context) {
	teamName := c.Param("team")
	_, err := services.DB.DeleteOne(teamsCollection, "name", teamName)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.AbortWithStatus(204)
}
