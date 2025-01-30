package domain

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Rotate(theta float64) Point {
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)

	return Point{
		X: p.X*cosTheta - p.Y*sinTheta,
		Y: p.X*sinTheta + p.Y*cosTheta,
	}
}
