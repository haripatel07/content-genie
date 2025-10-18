package main

import (
	"log"
	"content-genie/backend/api"    
	"content-genie/backend/config" 
	"content-genie/backend/models" 

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Load Configuration from .env file
	config.LoadConfig()

	// 2. Initialize Database
	db, err := gorm.Open(sqlite.Open("content_genie.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// 3. Auto-migrate the schema for our Job model
	err = db.AutoMigrate(&models.Job{})
	if err != nil {
		log.Fatal("Failed to migrate database schema")
	}

	// 4. Setup Gin Router
	r := gin.Default()

	// 5. Setup API routes, passing the database connection
	api.SetupRoutes(r, db)

	// 6. Start the server
	log.Println("Server starting on http://localhost:8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
