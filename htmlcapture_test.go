package htmlcapture

import (
	"testing"
)

func TestCapture(t *testing.T) {
	opts := CaptureOptions{
		Input: `<html><body><h1>Test</h1></body></html>`,
	}
	img, err := Capture(opts)
	if err != nil {
		t.Errorf("Capture failed: %v", err)
	}
	if len(img) == 0 {
		t.Error("Expected non-empty image bytes")
	}
}
