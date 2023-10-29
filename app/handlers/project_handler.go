package handlers

import (
	"net/http"

	"BackendService/app/models"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	siakod := models.Discipline{
		ID:   1,
		Name: "siakod",
	}
	projects := []models.Project{
		{
			ID:         "1",
			Name:       "ProjectHub",
			Annotation: "Some annot",
			UploadDate: "29.10.2023",
			Discipline: siakod,
		},
		{
			ID:         "1",
			Name:       "ProjectHub",
			Annotation: "Some annot",
			UploadDate: "29.10.2023",
			Discipline: siakod,
		},
	}
	c.JSON(http.StatusOK, projects)
}

func GetProject(c *gin.Context, db *sql.DB) {
	projectID := c.Param("id")

	query := "SELECT * FROM projects WHERE id = $1"
	row := db.QueryRow(query, projectID)

	var project models.Project
	err := row.Scan(&project.ID, &project.Name, &project.Annotation, &project.UploadDate, &project.Discipline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
