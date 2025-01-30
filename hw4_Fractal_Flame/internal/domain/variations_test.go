package domain_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

func TestLinearTransform(t *testing.T) {
	v := &domain.Linear{}
	p := domain.Point{X: 1.0, Y: 2.0}

	result := v.Transform(p)

	assert.Equal(t, p, result)
}

func TestSinusoidalTransform(t *testing.T) {
	v := &domain.Sinusoidal{}
	p := domain.Point{X: math.Pi / 2, Y: math.Pi}

	result := v.Transform(p)

	assert.InDelta(t, 1.0, result.X, 1e-6)
	assert.InDelta(t, 0.0, result.Y, 1e-6)
}

func TestSphericalTransform(t *testing.T) {
	v := &domain.Spherical{}
	p := domain.Point{X: 1.0, Y: 1.0}

	result := v.Transform(p)

	expected := domain.Point{
		X: 0.5,
		Y: 0.5,
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestSwirlTransform(t *testing.T) {
	v := &domain.Swirl{}
	p := domain.Point{X: 1.0, Y: 1.0}
	r2 := p.X*p.X + p.Y*p.Y

	result := v.Transform(p)

	expectedX := p.X*math.Sin(r2) - p.Y*math.Cos(r2)
	expectedY := p.X*math.Cos(r2) + p.Y*math.Sin(r2)
	assert.InDelta(t, expectedX, result.X, 1e-6)
	assert.InDelta(t, expectedY, result.Y, 1e-6)
}

func TestHorseshoeTransform(t *testing.T) {
	v := &domain.Horseshoe{}
	p := domain.Point{X: 1.0, Y: 1.0}
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)

	result := v.Transform(p)

	expected := domain.Point{
		X: (p.X - p.Y) * (p.X + p.Y) / r,
		Y: 2 * p.X * p.Y / r,
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestFisheyeTransform(t *testing.T) {
	v := &domain.Fisheye{}
	p := domain.Point{X: 1.0, Y: 0.0}
	r := 2 / (math.Sqrt(p.X*p.X+p.Y*p.Y) + 1)

	result := v.Transform(p)

	expected := domain.Point{
		X: r * p.Y,
		Y: r * p.X,
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestPolarTransform(t *testing.T) {
	v := &domain.Polar{}
	p := domain.Point{X: 1.0, Y: 1.0}
	theta := math.Atan2(p.Y, p.X)
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)

	result := v.Transform(p)

	expected := domain.Point{
		X: theta / math.Pi,
		Y: r - 1.0,
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestDiskTransform(t *testing.T) {
	v := &domain.Disk{}
	p := domain.Point{X: 1.0, Y: 0.0}
	r := math.Pi * math.Sqrt(p.X*p.X+p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	result := v.Transform(p)

	expected := domain.Point{
		X: (theta / math.Pi) * math.Sin(r),
		Y: (theta / math.Pi) * math.Cos(r),
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestBubbleTransform(t *testing.T) {
	v := &domain.Bubble{}
	p := domain.Point{X: 1.0, Y: 1.0}
	r2 := p.X*p.X + p.Y*p.Y
	scale := 4.0 / (r2 + 4.0)

	result := v.Transform(p)

	expected := domain.Point{
		X: p.X * scale,
		Y: p.Y * scale,
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestHeartTransform(t *testing.T) {
	v := &domain.Heart{}
	p := domain.Point{X: 1.0, Y: 1.0}
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	result := v.Transform(p)

	expected := domain.Point{
		X: r * math.Sin(theta*r),
		Y: -r * math.Cos(theta*r),
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestSpiralTransform(t *testing.T) {
	v := &domain.Spiral{}
	p := domain.Point{X: 1.0, Y: 1.0}
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	result := v.Transform(p)

	expected := domain.Point{
		X: (1 / r) * (math.Cos(theta) + math.Sin(r)),
		Y: (1 / r) * (math.Sin(theta) - math.Cos(r)),
	}
	assert.InDelta(t, expected.X, result.X, 1e-6)
	assert.InDelta(t, expected.Y, result.Y, 1e-6)
}

func TestChooseTransformation(t *testing.T) {
	transformations := domain.Transformations{
		{Weight: 1.0},
		{Weight: 2.0},
		{Weight: 3.0},
	}
	//nolint:gosec // не требуется сильная рандомизация
	rng := rand.New(rand.NewSource(42))

	counts := make([]int, len(transformations))
	iterations := 10000

	for i := 0; i < iterations; i++ {
		t := transformations.ChooseTransformation(rng)
		for idx, trans := range transformations {
			if t.Weight == trans.Weight {
				counts[idx]++
				break
			}
		}
	}

	assert.True(t, counts[1] > counts[0], "counts[1] должен быть больше counts[0]")
	assert.True(t, counts[2] > counts[1], "counts[2] должен быть больше counts[1]")
}

func TestNewVariation(t *testing.T) {
	t.Run("Позитивные случаи", func(t *testing.T) {
		validTests := []struct {
			name    string
			varType string
			want    interface{}
		}{
			{"linear", "linear", &domain.Linear{}},
			{"sinusoidal", "sinusoidal", &domain.Sinusoidal{}},
			{"spherical", "spherical", &domain.Spherical{}},
			{"swirl", "swirl", &domain.Swirl{}},
			{"horseshoe", "horseshoe", &domain.Horseshoe{}},
			{"fisheye", "fisheye", &domain.Fisheye{}},
			{"polar", "polar", &domain.Polar{}},
			{"disk", "disk", &domain.Disk{}},
			{"bubble", "bubble", &domain.Bubble{}},
			{"heart", "heart", &domain.Heart{}},
			{"spiral", "spiral", &domain.Spiral{}},
		}

		for _, tt := range validTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := domain.NewVariation(tt.varType)

				require.NoError(t, err)
				require.NotNil(t, got)

				require.IsType(t, tt.want, got, "Ожидался тип %T, но получил %T", tt.want, got)
			})
		}
	})

	t.Run("Негативные случаи", func(t *testing.T) {
		invalidTests := []struct {
			name    string
			varType string
		}{
			{"unknown", "unknown"},
		}

		for _, tt := range invalidTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := domain.NewVariation(tt.varType)

				require.Error(t, err)
				require.Nil(t, got)

				var unknownErr *domain.UnknownVariationError

				require.ErrorAs(t, err, &unknownErr)
				require.Equal(t, tt.varType, unknownErr.Name)
			})
		}
	})
}
