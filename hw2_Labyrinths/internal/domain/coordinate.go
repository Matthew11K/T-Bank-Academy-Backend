package domain

import "math"

type Coordinate struct {
	Row int
	Col int
}

func (c Coordinate) IsValid(height, width int) bool {
	return c.Row >= 0 && c.Row < height && c.Col >= 0 && c.Col < width
}

func (c Coordinate) Heuristic(other Coordinate) float64 {
	return math.Abs(float64(c.Row-other.Row)) + math.Abs(float64(c.Col-other.Col))
}
