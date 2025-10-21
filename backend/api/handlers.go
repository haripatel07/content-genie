package api

import (
	"content-genie/backend/models"
	"content-genie/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// JobRequest for creating a new job
type JobRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// CreateJobHandler handles the creation of a new job
func CreateJobHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json JobRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL provided."})
			return
		}

		job := models.Job{
			OriginalURL: json.URL,
			Status:      "pending",
		}

		result := db.Create(&job)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job in database."})
			return
		}

		// Trigger the processing in the background
		go services.ProcessJob(db, job.ID)

		c.JSON(http.StatusAccepted, job)
	}
}

// GetJobsHandler retrieves all jobs, newest first
func GetJobsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jobs []models.Job
		db.Order("created_at desc").Find(&jobs)
		c.JSON(http.StatusOK, jobs)
	}
}
