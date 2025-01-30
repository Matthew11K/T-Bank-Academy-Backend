package application_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/application/mocks"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/pkg/logger"
)

func BenchmarkSingleThreaded(b *testing.B) {
	config := application.Config{
		Width:           1024,
		Height:          768,
		Iterations:      10000000,
		Points:          100,
		Threads:         1,
		Output:          "benchmark_single.png",
		Gamma:           2.2,
		Symmetry:        1,
		Transformations: getBenchmarkTransformations(),
	}

	mockSaver := new(mocks.ImageSaver)
	mockSaver.On("Save", mock.Anything, "benchmark_single.png", 2.2).Return(nil)

	loggerInstance := logger.NewLogger(io.Discard)
	app := application.NewApplication(&config, mockSaver, loggerInstance)

	b.ResetTimer()

	start := time.Now()

	for i := 0; i < b.N; i++ {
		app.Render()
	}

	duration := time.Since(start)

	b.StopTimer()

	b.Logf("Однопоточная версия заняла: %v", duration)

	os.Remove("benchmark_single.png")
}

func BenchmarkRenderMultiThreaded_LessThanCores(b *testing.B) {
	config := application.Config{
		Width:           1024,
		Height:          768,
		Iterations:      10000000,
		Points:          100,
		Threads:         4,
		Output:          "benchmark_multi_less.png",
		Gamma:           2.2,
		Symmetry:        1,
		Transformations: getBenchmarkTransformations(),
	}

	mockSaver := new(mocks.ImageSaver)
	mockSaver.On("Save", mock.Anything, "benchmark_multi_less.png", 2.2).Return(nil)

	loggerInstance := logger.NewLogger(io.Discard)
	app := application.NewApplication(&config, mockSaver, loggerInstance)

	b.ResetTimer()

	start := time.Now()

	for i := 0; i < b.N; i++ {
		app.RenderMultiThreaded(config.Threads)
	}

	duration := time.Since(start)

	b.StopTimer()

	b.Logf("Многопоточная версия (< ядер) заняла: %v", duration)

	os.Remove("benchmark_multi_less.png")
}

func BenchmarkRenderMultiThreaded_EqualToCores(b *testing.B) {
	config := application.Config{
		Width:           1024,
		Height:          768,
		Iterations:      10000000,
		Points:          100,
		Threads:         8,
		Output:          "benchmark_multi_equal.png",
		Gamma:           2.2,
		Symmetry:        1,
		Transformations: getBenchmarkTransformations(),
	}

	mockSaver := new(mocks.ImageSaver)
	mockSaver.On("Save", mock.Anything, "benchmark_multi_equal.png", 2.2).Return(nil)

	loggerInstance := logger.NewLogger(io.Discard)
	app := application.NewApplication(&config, mockSaver, loggerInstance)

	b.ResetTimer()

	start := time.Now()

	for i := 0; i < b.N; i++ {
		app.RenderMultiThreaded(config.Threads)
	}

	duration := time.Since(start)

	b.StopTimer()

	b.Logf("Многопоточная версия (= ядер) заняла: %v", duration)

	os.Remove("benchmark_multi_equal.png")
}

func BenchmarkRenderMultiThreaded_MoreThanCores(b *testing.B) {
	config := application.Config{
		Width:           1024,
		Height:          768,
		Iterations:      10000000,
		Points:          100,
		Threads:         10,
		Output:          "benchmark_multi_more.png",
		Gamma:           2.2,
		Symmetry:        1,
		Transformations: getBenchmarkTransformations(),
	}

	mockSaver := new(mocks.ImageSaver)
	mockSaver.On("Save", mock.Anything, "benchmark_multi_more.png", 2.2).Return(nil)

	loggerInstance := logger.NewLogger(io.Discard)
	app := application.NewApplication(&config, mockSaver, loggerInstance)

	b.ResetTimer()

	start := time.Now()

	for i := 0; i < b.N; i++ {
		app.RenderMultiThreaded(config.Threads)
	}

	duration := time.Since(start)

	b.StopTimer()

	b.Logf("Многопоточная версия (> ядер) заняла: %v", duration)

	os.Remove("benchmark_multi_more.png")
}

func getBenchmarkTransformations() []domain.Transformation {
	transformations := []domain.Transformation{
		{
			Affine: domain.AffineTransform{
				A: 0.5, B: 0.0, C: 0.0,
				D: 0.0, E: 0.5, F: 0.0,
			},
			Variation: &domain.Linear{},
			Color:     [3]float64{1.0, 0.0, 0.0},
			Weight:    1.0,
		},
		{
			Affine: domain.AffineTransform{
				A: 0.5, B: 0.0, C: 1.0,
				D: 0.0, E: 0.5, F: 0.0,
			},
			Variation: &domain.Sinusoidal{},
			Color:     [3]float64{0.0, 1.0, 0.0},
			Weight:    1.0,
		},
	}

	return transformations
}
