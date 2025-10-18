package services

import (
	"content-genie/backend/models"
	"log"

	"gorm.io/gorm"
)

// ProcessJob orchestrates the scraping and AI content generation for a given job ID
func ProcessJob(db *gorm.DB, jobID uint) {
	log.Printf("Starting to process job ID: %d", jobID)
	var job models.Job
	if err := db.First(&job, jobID).Error; err != nil {
		log.Printf("Error: Could not find job ID %d to process. %v", jobID, err)
		return
	}

	// 1. Scrape the content from the URL
	job.Status = "processing"
	job.StatusDetail = "Scraping article content..."
	db.Save(&job)

	scrapedText, err := ScrapeArticleContent(job.OriginalURL)
	if err != nil {
		log.Printf("Error scraping URL for job %d: %v", jobID, err)
		job.Status = "failed"
		job.StatusDetail = "Failed to scrape content from URL."
		db.Save(&job)
		return
	}

	// 2. Generate content with AI
	job.StatusDetail = "Generating content with AI..."
	db.Save(&job)

	generatedContent, err := GenerateContentWithAI(scrapedText)
	if err != nil {
		log.Printf("Error generating AI content for job %d: %v", jobID, err)
		job.Status = "failed"
		job.StatusDetail = "Failed to generate content from AI."
		db.Save(&job)
		return
	}

	// 3. Update the job with the results and mark as complete
	job.Summary = generatedContent.Summary
	job.LinkedInPost = generatedContent.LinkedInPost
	if err := job.SetTweets(generatedContent.Tweets); err != nil {
		log.Printf("Error setting tweets for job %d: %v", jobID, err)
		job.Status = "failed"
		job.StatusDetail = "Failed to serialize generated tweets."
		db.Save(&job)
		return
	}

	job.Status = "complete"
	job.StatusDetail = ""
	db.Save(&job)

	log.Printf("Successfully processed job ID: %d", jobID)
}
