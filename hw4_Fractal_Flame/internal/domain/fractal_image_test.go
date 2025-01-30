package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

func TestFractalImageUpdatePixel(t *testing.T) {
	t.Run("Single update", func(t *testing.T) {
		fi := domain.NewFractalImage(100, 100)
		color := [3]float64{1.0, 0.0, 0.0}

		fi.UpdatePixel(50, 50, color)
		pixel := fi.GetPixel(50, 50)

		assert.Equal(t, 1, pixel.HitCount)
		assert.Equal(t, color[0], pixel.R)
		assert.Equal(t, color[1], pixel.G)
		assert.Equal(t, color[2], pixel.B)
	})

	t.Run("Multiple updates", func(t *testing.T) {
		fi := domain.NewFractalImage(100, 100)
		x, y := 50, 50

		color1 := [3]float64{1.0, 0.0, 0.0}

		fi.UpdatePixel(x, y, color1)
		pixel := fi.GetPixel(x, y)

		assert.Equal(t, 1, pixel.HitCount)
		assert.Equal(t, color1[0], pixel.R)
		assert.Equal(t, color1[1], pixel.G)
		assert.Equal(t, color1[2], pixel.B)

		color2 := [3]float64{0.0, 1.0, 0.0}

		fi.UpdatePixel(x, y, color2)
		pixel = fi.GetPixel(x, y)

		assert.Equal(t, 2, pixel.HitCount)
		assert.InDelta(t, 0.5, pixel.R, 1e-6)
		assert.InDelta(t, 0.5, pixel.G, 1e-6)
		assert.Equal(t, 0.0, pixel.B)
	})
}
func TestFractalImageContains(t *testing.T) {
	fi := domain.NewFractalImage(100, 100)

	testCases := []struct {
		name     string
		x, y     int
		expected bool
	}{
		{"В пределах", 50, 50, true},
		{"На границе", 99, 99, true},
		{"Вне X", 100, 50, false},
		{"Вне Y", 50, 100, false},
		{"Отрицательный X", -1, 50, false},
		{"Отрицательный Y", 50, -1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := fi.Contains(tc.x, tc.y)

			assert.Equal(t, tc.expected, result)
		})
	}
}
