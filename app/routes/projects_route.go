package routes

import (
	"BackendService/app/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/projects", handlers.GetProjects)

	router.GET("/projects/:id", func(c *gin.Context) {
		handlers.GetProject(c, db)
	})

	return router
}
