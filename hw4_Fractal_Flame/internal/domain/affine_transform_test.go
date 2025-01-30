package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

func TestAffineTransformApply(t *testing.T) {
	at := domain.AffineTransform{
		A: 1.0, B: 0.0, C: 2.0,
		D: 0.0, E: 1.0, F: 3.0,
	}
	p := domain.Point{X: 1.0, Y: 1.0}

	result := at.Apply(p)

	expected := domain.Point{X: 3.0, Y: 4.0}
	assert.Equal(t, expected, result)
}
