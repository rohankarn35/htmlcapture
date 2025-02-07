package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rohankarn35/htmlcapture"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>IPO Alert</title>
    <style>
      body {
        font-family: "Poppins", sans-serif;
        text-align: center;
        background: linear-gradient(135deg, #f3f4f6, #ffffff);
        padding: 20px;
      }
      .container {
        background: white;
        max-width: 500px;
        margin: auto;
        padding: 25px;
        border-radius: 15px;
        box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.1);
      }
      .logo {
        font-size: 18px;
        font-weight: bold;
        color: #b72b67;
      }
      .title {
        font-size: 28px;
        font-weight: bold;
        margin-top: 10px;
        color: #333;
      }
      .subtitle {
        color: green;
        font-size: 18px;
        font-weight: bold;
      }
      .highlight {
        background: #ffcc00;
        padding: 8px 12px;
        font-weight: bold;
        border-radius: 8px;
        display: inline-block;
        margin-top: 15px;
        font-size: 16px;
      }
      .details {
        margin-top: 20px;
        border: 1px solid #ddd;
        border-radius: 10px;
        padding: 15px;
        text-align: left;
        background: #f9f9f9;
      }
      .details table {
        width: 100%;
        border-collapse: collapse;
      }
      .details td {
        padding: 10px;
        border-bottom: 1px solid #eee;
        font-size: 16px;
      }
      .details tr:last-child td {
        border-bottom: none;
      }
      .disclaimer {
        font-size: 13px;
        margin-top: 15px;
        color: #555;
      }
      .footer {
        margin-top: 25px;
        font-size: 15px;
        color: #007a33;
        font-weight: bold;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="logo">{{.CompanyName}}</div>
      <div class="title">{{.Title}}</div>
      <div class="subtitle">{{.Subtitle}}</div>

      <div class="details">
        <table>
          <tr>
            <td><strong>Issue Date</strong></td>
            <td>{{.IssueDate}}</td>
          </tr>
          <tr>
            <td><strong>Closing Date</strong></td>
            <td>{{.ClosingDate}}</td>
          </tr>
          <tr>
            <td><strong>Issue Price</strong></td>
            <td>{{.IssuePrice}}</td>
          </tr>
          <tr>
            <td><strong>Sector</strong></td>
            <td>{{.Sector}}</td>
          </tr>
        </table>
      </div>

      <div class="footer">NEPSE Navigator<br />t.me/nepsenavigator</div>
    </div>
  </body>
</html>
`

func main() {
	// Define dynamic variables

	// Prepare capture options
	opts := htmlcapture.CaptureOptions{
		Input: htmlTemplate,
		Variables: map[string]string{
			"CompanyName": "ABC Corp",
			"Title":       "IPO Alert",
			"Subtitle":    "(For Equity)",

			"IssueDate":   "2023-10-01",
			"ClosingDate": "2023-10-10",
			"IssuePrice":  "Rs. 100",
			"Sector":      "Technology",
		},
		Selector: ".container",
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
