package domain

type FractalImage struct {
	Data   [][]Pixel
	Width  int
	Height int
}

func NewFractalImage(width, height int) *FractalImage {
	data := make([][]Pixel, height)

	for i := range data {
		data[i] = make([]Pixel, width)
	}

	return &FractalImage{
		Data:   data,
		Width:  width,
		Height: height,
	}
}

func (fi *FractalImage) Contains(x, y int) bool {
	return x >= 0 && x < fi.Width && y >= 0 && y < fi.Height
}

func (fi *FractalImage) UpdatePixel(x, y int, color [3]float64) {
	pixel := &fi.Data[y][x]
	oldCount := pixel.HitCount
	pixel.HitCount++

	if oldCount == 0 {
		pixel.R = color[0]
		pixel.G = color[1]
		pixel.B = color[2]
	} else {
		weight := 1.0 / float64(pixel.HitCount)
		oldWeight := float64(oldCount) / float64(pixel.HitCount)

		pixel.R = pixel.R*oldWeight + color[0]*weight
		pixel.G = pixel.G*oldWeight + color[1]*weight
		pixel.B = pixel.B*oldWeight + color[2]*weight
	}
}
func (fi *FractalImage) GetPixel(x, y int) Pixel {
	return fi.Data[y][x]
}
