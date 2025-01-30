package application

import (
	"math"
	"math/rand"
	"sync"

	"golang.org/x/exp/slog"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

type ImageSaver interface {
	Save(image *domain.FractalImage, filename string, gamma float64) error
}

type Application struct {
	config     *Config
	imageSaver ImageSaver
	logger     *slog.Logger
	canvas     *domain.FractalImage
	world      domain.Rect
}

func NewApplication(config *Config, imageSaver ImageSaver, logger *slog.Logger) *Application {
	return &Application{
		config:     config,
		imageSaver: imageSaver,
		logger:     logger,
		canvas:     domain.NewFractalImage(config.Width, config.Height),
		world: domain.Rect{
			X:      -1.0,
			Y:      -1.0,
			Width:  2.0,
			Height: 2.0,
		},
	}
}

func (app *Application) Run() error {
	app.logger.Info("Начало генерации", slog.Any("config", app.config))

	if app.config.Threads <= 1 {
		app.logger.Info("Запуск в однопоточке")
		app.Render()
	} else {
		app.logger.Info("Запуск в многопоточке",
			slog.Int("threads", app.config.Threads))
		app.RenderMultiThreaded(app.config.Threads)
	}

	if err := app.imageSaver.Save(app.canvas, app.config.Output,
		app.config.Gamma); err != nil {
		return WrapError(err, "не удалось сохранить изображение")
	}

	return nil
}

func (app *Application) Render() {
	for range app.config.Points {
		app.renderIterations()
	}
}

func (app *Application) renderIterations() {
	//nolint:gosec // не требуется сильная рандомизация
	rng := rand.New(rand.NewSource(rand.Int63()))

	point := app.world.RandomPoint(rng)
	burnIn := 20
	totalIterations := app.config.Iterations + burnIn

	for i := 0; i < totalIterations; i++ {
		transformation := app.config.Transformations.ChooseTransformation(rng)

		point = transformation.Affine.Apply(point)
		point = transformation.Variation.Transform(point)

		if i >= burnIn {
			x := int((point.X - app.world.X) / app.world.Width *
				float64(app.canvas.Width))
			y := int((point.Y - app.world.Y) / app.world.Height *
				float64(app.canvas.Height))

			if app.canvas.Contains(x, y) {
				app.canvas.UpdatePixel(x, y, transformation.Color)
			}

			for s := 1; s < app.config.Symmetry; s++ {
				theta := 2 * math.Pi * float64(s) /
					float64(app.config.Symmetry)
				rotated := point.Rotate(theta)
				rx := int((rotated.X - app.world.X) / app.world.Width *
					float64(app.canvas.Width))
				ry := int((rotated.Y - app.world.Y) / app.world.Height *
					float64(app.canvas.Height))

				if app.canvas.Contains(rx, ry) {
					app.canvas.UpdatePixel(rx, ry, transformation.Color)
				}
			}
		}
	}
}

func (app *Application) RenderMultiThreaded(threads int) {
	var wg sync.WaitGroup

	pointsPerThread := app.config.Points / threads
	extraPoints := app.config.Points % threads

	segmentHeight := app.canvas.Height / threads
	extraSegment := app.canvas.Height % threads

	for i := 0; i < threads; i++ {
		wg.Add(1)

		yStart := i*segmentHeight + min(i, extraSegment)
		yEnd := yStart + segmentHeight

		if i < extraSegment {
			yEnd++
		}

		pointsForThisThread := pointsPerThread
		if i < extraPoints {
			pointsForThisThread++
		}

		go func() {
			defer wg.Done()
			app.renderIterationsSegment(yStart, yEnd, pointsForThisThread)
		}()
	}

	wg.Wait()
}

func (app *Application) renderIterationsSegment(yStart, yEnd, points int) {
	//nolint:gosec // не требуется сильная рандомизация
	rng := rand.New(rand.NewSource(rand.Int63()))

	for p := 0; p < points; p++ {
		point := app.world.RandomPoint(rng)
		burnIn := 20
		totalIterations := app.config.Iterations + burnIn

		for i := 0; i < totalIterations; i++ {
			transformation := app.config.Transformations.ChooseTransformation(rng)
			point = transformation.Affine.Apply(point)
			point = transformation.Variation.Transform(point)

			if i >= burnIn && app.world.Contains(point) {
				x := int((point.X - app.world.X) / app.world.Width * float64(app.canvas.Width))
				y := int((point.Y - app.world.Y) / app.world.Height * float64(app.canvas.Height))

				if y >= yStart && y < yEnd && app.canvas.Contains(x, y) {
					app.canvas.UpdatePixel(x, y, transformation.Color)
				}

				for s := 1; s < app.config.Symmetry; s++ {
					theta := 2 * math.Pi * float64(s) / float64(app.config.Symmetry)
					rotated := point.Rotate(theta)
					rx := int((rotated.X - app.world.X) / app.world.Width * float64(app.canvas.Width))
					ry := int((rotated.Y - app.world.Y) / app.world.Height * float64(app.canvas.Height))

					if ry >= yStart && ry < yEnd && app.canvas.Contains(rx, ry) {
						app.canvas.UpdatePixel(rx, ry, transformation.Color)
					}
				}
			}
		}
	}
}
