package infrastructure

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"
)

type ImageSaverImpl struct{}

func (s *ImageSaverImpl) Save(fi *domain.FractalImage, filename string, gamma float64) error {
	img := image.NewRGBA(image.Rect(0, 0, fi.Width, fi.Height))
	maxLogHitCount := 0.0

	logValues := make([][]float64, fi.Height)
	for y := 0; y < fi.Height; y++ {
		logValues[y] = make([]float64, fi.Width)

		for x := 0; x < fi.Width; x++ {
			pixel := fi.GetPixel(x, y)
			if pixel.HitCount > 0 {
				logHit := math.Log10(float64(pixel.HitCount))
				logValues[y][x] = logHit

				if logHit > maxLogHitCount {
					maxLogHitCount = logHit
				}
			}
		}
	}

	for y := 0; y < fi.Height; y++ {
		for x := 0; x < fi.Width; x++ {
			pixel := fi.GetPixel(x, y)
			if pixel.HitCount > 0 {
				normalized := logValues[y][x] / maxLogHitCount
				scale := math.Pow(normalized, 1.0/gamma)

				r := uint8(math.Min(255, pixel.R*scale*255))
				g := uint8(math.Min(255, pixel.G*scale*255))
				b := uint8(math.Min(255, pixel.B*scale*255))

				img.Set(x, y, color.NRGBA{
					R: r, G: g, B: b, A: 255})
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
