package controllers

import (
	"GID/helpers"
	"GID/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(helpers.Log), gin.Recovery()) //Setup logging and panic recovery
	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           72 * time.Hour,

		// Check with
		/*cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Authorization"},
		},*/
	}))

	// API routes
	apiGroup := router.Group("/api").Use(middleware.EnsureValidToken)
	{
		boardsController := BoardsController{}
		apiGroup.GET("/boards", boardsController.GetAllBoards)
		apiGroup.POST("/boards", boardsController.CreateBoard)
		apiGroup.GET("/boards/:boardId", boardsController.GetBoard)
		apiGroup.DELETE("/boards/:boardId", boardsController.DeleteBoard)
		apiGroup.GET("/boards/:boardId/backlog", boardsController.GetBacklog)
		apiGroup.GET("/boards/:boardId/issues", boardsController.GetIssues)

		issuesController := IssuesController{}
		apiGroup.POST("/issues", issuesController.CreateIssue)
		apiGroup.GET("/issues/:issueKeyOrId", issuesController.GetIssue)
		apiGroup.PUT("/issues/:issueKeyOrId", issuesController.UpdateIssue)
		apiGroup.DELETE("/issues/:issueKeyOrId", issuesController.DeleteIssue)

		sessionsController := SessionsController{}
		apiGroup.GET("/sessions", sessionsController.GetCurrentUser)
		apiGroup.POST("/sessions", sessionsController.Login)
		apiGroup.DELETE("/sessions", sessionsController.Logout)

		// Projects endpoints
		projectsController := ProjectsController{}
		apiGroup.POST("/projects", projectsController.CreateProject)
		apiGroup.GET("/projects", projectsController.GetAllProjects)
		apiGroup.GET("/projects/:projectIdOrKey", projectsController.GetDescription)
		apiGroup.PUT("/projects/:projectIdOrKey", projectsController.UpdateProject)
		apiGroup.DELETE("/projects/:projectIdOrKey", projectsController.DeleteProject)

		teamsController := TeamsController{}
		apiGroup.POST("/teams", teamsController.CreateTeam)
		apiGroup.GET("/teams/:teamId", teamsController.GetMembers)
		apiGroup.POST("/teams/:teamId", teamsController.UpdateTeam)
		apiGroup.DELETE("/teams/:teamId", teamsController.DeleteTeam)
	}
	return router
}
