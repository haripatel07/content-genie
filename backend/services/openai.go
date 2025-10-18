package services

import (
	"content-genie/backend/config"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// GeneratedContent holds all the pieces of content created by the AI
type GeneratedContent struct {
	Summary      string
	Tweets       []string
	LinkedInPost string
}

// GenerateContentWithAI sends the scraped text to OpenAI and asks for repurposed content
func GenerateContentWithAI(scrapedText string) (*GeneratedContent, error) {
	client := openai.NewClient(config.AppConfig.OpenAIAPIKey)

	// A single, complex prompt is more efficient than multiple API calls
	prompt := fmt.Sprintf(`
        Based on the following article text, please generate three distinct pieces of content:
        1. A concise summary of the article (around 100 words).
        2. Three engaging tweets, each under 280 characters.
        3. A professional LinkedIn post (around 150 words).

        Format the output EXACTLY as follows, using the specified separators:

        [SUMMARY]
        {Your summary here}
        [END_SUMMARY]

        [TWEETS]
        1. {First tweet here}
        2. {Second tweet here}
        3. {Third tweet here}
        [END_TWEETS]

        [LINKEDIN]
        {Your LinkedIn post here}
        [END_LINKEDIN]

        ---
        ARTICLE TEXT:
        %s
    `, scrapedText)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("OpenAI chat completion error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("OpenAI returned no choices")
	}

	// Parse the structured response
	return parseAIResponse(resp.Choices[0].Message.Content)
}

// parseAIResponse extracts content from the structured string returned by the AI
func parseAIResponse(response string) (*GeneratedContent, error) {
	summary, err := extractContent(response, "[SUMMARY]", "[END_SUMMARY]")
	if err != nil {
		return nil, fmt.Errorf("failed to parse summary: %w", err)
	}

	tweetsBlock, err := extractContent(response, "[TWEETS]", "[END_TWEETS]")
	if err != nil {
		return nil, fmt.Errorf("failed to parse tweets: %w", err)
	}

	linkedinPost, err := extractContent(response, "[LINKEDIN]", "[END_LINKEDIN]")
	if err != nil {
		return nil, fmt.Errorf("failed to parse LinkedIn post: %w", err)
	}

	// Split tweets by newline
	var tweets []string
	rawTweets := strings.Split(tweetsBlock, "\n")
	for _, t := range rawTweets {
		trimmed := strings.TrimSpace(t)
		// Remove numbering like "1. ", "2. ", etc.
		if len(trimmed) > 3 && trimmed[1] == '.' {
			trimmed = trimmed[3:]
		}
		if trimmed != "" {
			tweets = append(tweets, trimmed)
		}
	}

	return &GeneratedContent{
		Summary:      summary,
		Tweets:       tweets,
		LinkedInPost: linkedinPost,
	}, nil
}

// Helper to extract content between two tags
func extractContent(text, startTag, endTag string) (string, error) {
	start := strings.Index(text, startTag)
	if start == -1 {
		return "", fmt.Errorf("start tag '%s' not found", startTag)
	}
	start += len(startTag)

	end := strings.Index(text[start:], endTag)
	if end == -1 {
		return "", fmt.Errorf("end tag '%s' not found", endTag)
	}

	return strings.TrimSpace(text[start : start+end]), nil
}
