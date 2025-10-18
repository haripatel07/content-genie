package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures the application's routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Configure CORS to allow frontend requests
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Your frontend URL
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.POST("/jobs", CreateJobHandler(db))
		api.GET("/jobs", GetJobsHandler(db))
	}
}
