package hillchart_test

import (
	"fmt"
	"image"
	"os"
	"reflect"
	"testing"

	"github.com/alex-evans/braglog/internal/hillchart"
)

func TestGenerateHillChart(t *testing.T) {
	tests := []struct {
		pointPercentage float64
		pointLabel      string
	}{
		{40, "TestPoint1"},
		{75, "TestPoint2"},
	}

	for _, tt := range tests {
		t.Run(tt.pointLabel, func(t *testing.T) {
			generatedImagePath := "normal.png"
			expectedImagePath := fmt.Sprintf("testImages/expected_%.0f.png", tt.pointPercentage)

			err := hillchart.GenerateHillChart(tt.pointPercentage, tt.pointLabel)
			if err != nil {
				t.Fatalf("GenerateHillChart(%f, %s) error: %v", tt.pointPercentage, tt.pointLabel, err)
			}

			expectedImageFile, err := os.Open(expectedImagePath)
			if err != nil {
				t.Fatalf("Error opening expected PNG: %v", err)
			}

			generatedImageFile, err := os.Open(generatedImagePath)
			if err != nil {
				t.Fatalf("Error opening generated PNG: %v", err)
			}

			expectedImage, _, err := image.Decode(expectedImageFile)
			if err != nil {
				t.Fatalf("Error decoding expected PNG: %v", err)
			}

			generatedImage, _, err := image.Decode(generatedImageFile)
			if err != nil {
				t.Fatalf("Error decoding generated PNG: %v", err)
			}

			if !reflect.DeepEqual(expectedImage, generatedImage) {
				t.Error("Generated image does not match the expected image.")
			}

		})
	}
}
