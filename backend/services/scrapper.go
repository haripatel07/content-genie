package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeArticleContent fetches and extracts the main text content from a URL.
func ScrapeArticleContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Try to find the main content, common in article tags
	// This is a simplistic approach and may need refinement for different sites
	var content strings.Builder
	doc.Find("article p, .post-content p, .entry-content p").Each(func(i int, s *goquery.Selection) {
		content.WriteString(s.Text())
		content.WriteString("\n\n")
	})

	textContent := content.String()

	if len(textContent) < 200 { // Fallback if the specific selectors failed
		textContent = doc.Find("body").Text()
	}

	// Basic cleanup
	cleanedText := strings.TrimSpace(textContent)
	if len(cleanedText) == 0 {
		return "", fmt.Errorf("could not extract meaningful content from the page")
	}

	return cleanedText, nil
}
