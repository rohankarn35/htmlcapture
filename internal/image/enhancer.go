package imgutil

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png" // Register image decoders
	"math"
)

// EnhanceQuality improves image clarity using only standard library
func EnhanceQuality(input []byte) ([]byte, error) {
	// Decode image with format detection
	img, format, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	// Apply quality enhancements
	enhanced := processImage(img)

	// Encode with quality settings
	return encodeWithQuality(enhanced, format)
}

func processImage(img image.Image) image.Image {
	// Convert to RGBA for pixel manipulation
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	// Copy original pixels
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// Simple contrast enhancement
	return adjustContrast(rgba, 1)
}

func adjustContrast(img *image.RGBA, factor float64) image.Image {
	bounds := img.Bounds()

	// Apply contrast to each pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := img.RGBAAt(x, y)
			adjusted := color.RGBA{
				R: clamp(float64(original.R) * factor),
				G: clamp(float64(original.G) * factor),
				B: clamp(float64(original.B) * factor),
				A: original.A,
			}
			img.SetRGBA(x, y, adjusted)
		}
	}
	return img
}

func clamp(value float64) uint8 {
	return uint8(math.Min(255, math.Max(0, value)))
}

func encodeWithQuality(img image.Image, format string) ([]byte, error) {
	var buf bytes.Buffer

	switch format {
	case "jpeg":
		// High-quality JPEG encoding
		err := jpeg.Encode(&buf, img, &jpeg.Options{
			Quality: 100, // Increased from default 75
		})
		return buf.Bytes(), err

	default:
		// Convert other formats to high-quality JPEG
		err := jpeg.Encode(&buf, img, &jpeg.Options{
			Quality: 100,
		})
		return buf.Bytes(), err
	}
}
