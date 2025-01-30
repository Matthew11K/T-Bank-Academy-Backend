package domain

import "math/rand"

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (r *Rect) Contains(p Point) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}

func (r Rect) RandomPoint(rng *rand.Rand) Point {
	x := r.X + rng.Float64()*r.Width
	y := r.Y + rng.Float64()*r.Height

	return Point{X: x, Y: y}
}
