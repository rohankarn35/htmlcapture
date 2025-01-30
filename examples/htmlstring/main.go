package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rohankarn35/htmlcapture"
)

func main() {
	// Define dynamic HTML content with placeholders
	htmlString := `<html><body><h1>Hello {{.Name}}, Welcome to {{.Place}}!</h1></body></html>`

	// Define variables to inject into the HTML content
	variables := map[string]string{
		"Name":  "Alice",
		"Place": "Wonderland",
	}

	// Prepare capture options
	opts := htmlcapture.CaptureOptions{
		Input:     htmlString, // Pass raw HTML string
		Variables: variables,  // Inject dynamic variables
		ViewportW: 1080,       // Optional: Default width (Instagram post size)
		ViewportH: 1350,       // Optional: Default height (Instagram post size)
		// Selector:  "h1",       //optinal : if not provided it will capture the fullscreenshot
	}

	// Capture the screenshot
	imageData, err := htmlcapture.Capture(opts)
	if err != nil {
		log.Fatalf("Error capturing screenshot: %v", err)
	} else {
		fmt.Println("Screenshot captured successfully!")

		// Ensure the file path exists before saving the image
		filePath := "screenshot.png"
		err = os.WriteFile(filePath, imageData, 0644)
		if err != nil {
			log.Fatalf("Failed to save screenshot: %v", err)
		} else {
			fmt.Printf("Screenshot saved successfully to %s\n", filePath)
		}
	}
}
