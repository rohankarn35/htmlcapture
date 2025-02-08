package htmlcapture

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/rohankarn35/htmlcapture/internal/browser"
	imgutil "github.com/rohankarn35/htmlcapture/internal/image"
)

type CaptureOptions struct {
	Input     string            // Can be a URL, HTML string, or file path
	ViewportW int               // Viewport width (default: 1080 for Instagram)
	ViewportH int               // Viewport height (default: 1350 for Instagram)
	Selector  string            // CSS selector to capture a specific element
	Variables map[string]string // Variables for dynamic HTML
}

var DefaultViewport = CaptureOptions{
	ViewportW: 375,
	ViewportH: 667,
}

func detectInputType(input string) (isFile bool, isURL bool) {
	// Check if the input is a valid URL
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		return false, true
	}

	// Check if the input ends with .html and if the file exists
	if strings.HasSuffix(input, ".html") {

		return true, false

	}

	// Otherwise, assume it's a raw HTML string
	return false, false
}

func processDynamicHTML(htmlContent string, variables map[string]string) (string, error) {
	tmpl, err := template.New("html").Parse(htmlContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, variables)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func Capture(opts CaptureOptions) ([]byte, error) {
	if opts.Input == "" {
		return nil, fmt.Errorf("input cannot be empty")
	}

	isFile, isURL := detectInputType(opts.Input)

	var htmlContent string
	if isFile {
		if _, err := os.Stat(opts.Input); err != nil {
			return nil, fmt.Errorf("invalid file path")
		}
		content, err := os.ReadFile(opts.Input) // Read file content
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}
		htmlContent = string(content)
	} else if !isURL {
		htmlContent = opts.Input
	}

	// ✅ Fix: Ensure file-based HTML is processed for dynamic content
	if len(opts.Variables) > 0 {
		processedHTML, err := processDynamicHTML(htmlContent, opts.Variables)
		if err != nil {
			return nil, fmt.Errorf("failed to process dynamic HTML: %w", err)
		}
		htmlContent = processedHTML
	}

	if opts.ViewportW == 0 {
		opts.ViewportW = DefaultViewport.ViewportW
	}
	if opts.ViewportH == 0 {
		opts.ViewportH = DefaultViewport.ViewportH
	}

	targetInput := opts.Input
	if isFile || !isURL {
		targetInput = "data:text/html;charset=utf-8," + htmlContent
	}

	imageData, err := browser.CaptureScreenshot(targetInput, isFile, opts.ViewportW, opts.ViewportH, opts.Selector)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}
	img, err := imgutil.EnhanceQuality(imageData)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %w", err)
	}

	return img, nil
}
