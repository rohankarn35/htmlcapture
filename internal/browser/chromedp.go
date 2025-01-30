package browser

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// CaptureScreenshot takes a screenshot of a URL, HTML string, or file.
func CaptureScreenshot(input string, isFile bool, viewportW, viewportH int, selector string) ([]byte, error) {
	// Create ChromeDP context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set a timeout to prevent long-running sessions
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string
	var targetURL string
	var waitForRender chromedp.Action

	// Handle different input types
	if isValidURL(input) {
		// If input is a URL, set it directly and wait for rendering
		targetURL = input
		waitForRender = chromedp.WaitReady("body") // Wait until the body is fully loaded
	} else {
		// If input is an HTML string, use it directly
		htmlContent = input
	}

	// If it's a file or an HTML string, use a data URL
	if isFile || (!isFile && !isValidURL(input)) {
		// Convert HTML content into a proper data URL
		targetURL = encodeHTML(htmlContent)
		waitForRender = chromedp.Sleep(2 * time.Second) // Wait a bit to let JavaScript execute
	}

	// Allocate buffer for the screenshot
	var buf []byte

	// Define ChromeDP actions
	tasks := chromedp.Tasks{
		chromedp.Navigate(targetURL),                                 // Navigate to the page
		chromedp.EmulateViewport(int64(viewportW), int64(viewportH)), // Set viewport size
		waitForRender, // Ensure page rendering is complete
	}

	// Capture the screenshot
	if selector != "" {
		tasks = append(tasks, chromedp.Screenshot(selector, &buf, chromedp.NodeVisible))
	} else {
		tasks = append(tasks, chromedp.FullScreenshot(&buf, 100)) // Full-page screenshot
	}

	// Run ChromeDP tasks
	if err := chromedp.Run(ctx, tasks...); err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}

	return buf, nil
}

// isValidURL checks if a string is a valid URL.
func isValidURL(s string) bool {
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}

// encodeHTML ensures HTML content is properly formatted for data URL usage.
func encodeHTML(html string) string {
	// Escape problematic characters
	return strings.ReplaceAll(html, "#", "%23")
}
