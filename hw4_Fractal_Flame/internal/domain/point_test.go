package domain_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

func TestRotate(t *testing.T) {
	p := domain.Point{X: 1.0, Y: 0.0}
	theta := math.Pi / 2

	result := p.Rotate(theta)

	assert.InDelta(t, 0.0, result.X, 1e-6)
	assert.InDelta(t, 1.0, result.Y, 1e-6)
}

func TestRandomPointInRect(t *testing.T) {
	rect := domain.Rect{
		X:      -1.0,
		Y:      -1.0,
		Width:  2.0,
		Height: 2.0,
	}
	//nolint:gosec // не требуется сильная рандомизация
	rng := rand.New(rand.NewSource(42))

	for i := 0; i < 1000; i++ {
		point := rect.RandomPoint(rng)

		assert.True(t, rect.Contains(point), "Точка должна находиться внутри прямоугольника")
		assert.True(t, point.X >= rect.X && point.X <= rect.X+rect.Width,
			"X координата должна быть в пределах прямоугольника")
		assert.True(t, point.Y >= rect.Y && point.Y <= rect.Y+rect.Height,
			"Y координата должна быть в пределах прямоугольника")
	}
}

func TestRectContains(t *testing.T) {
	rect := domain.Rect{
		X:      0.0,
		Y:      0.0,
		Width:  2.0,
		Height: 2.0,
	}

	testCases := []struct {
		name     string
		point    domain.Point
		expected bool
	}{
		{"Внутри", domain.Point{X: 1.0, Y: 1.0}, true},
		{"На границе", domain.Point{X: 2.0, Y: 2.0}, true},
		{"Вне X", domain.Point{X: 2.1, Y: 1.0}, false},
		{"Вне Y", domain.Point{X: 1.0, Y: 2.1}, false},
		{"Вне обоих", domain.Point{X: -0.1, Y: -0.1}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := rect.Contains(tc.point)

			assert.Equal(t, tc.expected, result)
		})
	}
}
