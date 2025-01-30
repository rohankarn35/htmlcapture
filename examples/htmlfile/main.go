package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rohankarn35/htmlcapture"
)

func main() {
	// Define dynamic variables
	variables := map[string]string{
		"Title": "Mr. X",
		"User":  "John Doe",
	}

	// Prepare capture options
	opts := htmlcapture.CaptureOptions{
		Input:     "index.html", // Path to the HTML file
		Variables: variables,    // Inject dynamic variables
	}

	// Capture the screenshot
	imageData, err := htmlcapture.Capture(opts)
	if err != nil {
		log.Fatalf("Error capturing screenshot: %v", err)
	} else {
		fmt.Println("Screenshot captured successfully!")
		// You can save or further process the imageData as needed
		// Example: Save the image to a file
		_ = os.WriteFile("screenshot_from_file.png", imageData, 0644)
	}
}
