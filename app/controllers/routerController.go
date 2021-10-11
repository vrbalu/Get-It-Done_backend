package controllers

import (
	"GIT/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	// Go routines for websockets
	router.Use(ginlogrus.Logger(helpers.Log), gin.Recovery()) //Setup logging and panic recovery
	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           72 * time.Hour,
	}))

	// API routes
	apiGroup := router.Group("/api")
	{
		boardsController := BoardsController{}
		boardsGroup := apiGroup.Group("/boards")
		{
			boardsGroup.GET("", boardsController.GetAllBoards)
			boardsGroup.POST("", boardsController.CreateBoard)
			boardsGroup.GET("/:boardId", boardsController.GetBoard)
			boardsGroup.DELETE("/:boardId", boardsController.DeleteBoard)
			boardsGroup.GET("/:boardId/backlog", boardsController.GetBacklog)
			boardsGroup.GET("/:boardId/issues", boardsController.GetIssues)

		}
		issuesController := IssuesController{}
		issuesGroup := apiGroup.Group("/issues")
		{
			issuesGroup.POST("", issuesController.CreateIssue)
			issuesGroup.GET("/:issueKeyOrId", issuesController.GetIssue)
			issuesGroup.PUT("/:issueKeyOrId", issuesController.UpdateIssue)
			issuesGroup.DELETE("/:issueKeyOrId", issuesController.DeleteIssue)
		}
		sessionsController := SessionsController{}
		sessionsGroup := apiGroup.Group("/sessions")
		{
			sessionsGroup.GET("", sessionsController.GetCurrentUser)
			sessionsGroup.POST("", sessionsController.Login)
			sessionsGroup.DELETE("", sessionsController.Logout)
		}
		// Projects endpoints
		projectsController := ProjectsController{}
		projectsGroup := apiGroup.Group("/projects")
		{
			projectsGroup.POST("", projectsController.CreateProject)
			projectsGroup.GET("", projectsController.GetAllProjects)
			projectsGroup.PUT("/:projectIdOrKey", projectsController.GetDescription)
			projectsGroup.GET("/:projectIdOrKey", projectsController.UpdateProject)
			projectsGroup.DELETE("/:projectIdOrKey", projectsController.DeleteProject)
		}
		teamsController := TeamsController{}
		teamsGroup := apiGroup.Group("/teams/:teamId")
		{
			teamsGroup.GET("", teamsController.GetMembers)
			teamsGroup.POST("", teamsController.UpdateTeam)
			teamsGroup.PUT("/:projectIdOrKey", teamsController.DeleteTeam)
		}
	}
	return router
}
