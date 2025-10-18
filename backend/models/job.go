package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// Job represents a single content repurposing task
type Job struct {
	gorm.Model
	OriginalURL  string `json:"original_url"`
	Status       string `json:"status" gorm:"index"` // "pending", "processing", "complete", "failed"
	StatusDetail string `json:"status_detail"`       // For storing error messages
	Summary      string `json:"summary" gorm:"type:text"`
	Tweets       string `json:"tweets" gorm:"type:text"` // A JSON array of strings
	LinkedInPost string `json:"linkedin_post" gorm:"type:text"`
}

// SetTweets safely marshals a slice of strings into a JSON string for storage
func (j *Job) SetTweets(tweets []string) error {
	bytes, err := json.Marshal(tweets)
	if err != nil {
		return err
	}
	j.Tweets = string(bytes)
	return nil
}

// GetTweets safely unmarshals the JSON string back into a slice of strings
func (j *Job) GetTweets() ([]string, error) {
	var tweets []string
	if j.Tweets == "" {
		return tweets, nil
	}
	err := json.Unmarshal([]byte(j.Tweets), &tweets)
	return tweets, err
}
