package domain

import (
	"math"
	"math/rand"
)

type Transformer interface {
	Transform(p Point) Point
}

type UnknownVariationError struct {
	Name string
}

func (e *UnknownVariationError) Error() string {
	return "неизвестная вариация: " + e.Name
}

type VariationType string

const (
	LinearVariation     VariationType = "linear"
	SinusoidalVariation VariationType = "sinusoidal"
	SphericalVariation  VariationType = "spherical"
	SwirlVariation      VariationType = "swirl"
	HorseshoeVariation  VariationType = "horseshoe"
	FisheyeVariation    VariationType = "fisheye"
	PolarVariation      VariationType = "polar"
	DiskVariation       VariationType = "disk"
	BubbleVariation     VariationType = "bubble"
	HeartVariation      VariationType = "heart"
	SpiralVariation     VariationType = "spiral"
)

func NewVariation(name string) (Transformer, error) {
	switch VariationType(name) {
	case LinearVariation:
		return &Linear{}, nil
	case SinusoidalVariation:
		return &Sinusoidal{}, nil
	case SphericalVariation:
		return &Spherical{}, nil
	case SwirlVariation:
		return &Swirl{}, nil
	case HorseshoeVariation:
		return &Horseshoe{}, nil
	case FisheyeVariation:
		return &Fisheye{}, nil
	case PolarVariation:
		return &Polar{}, nil
	case DiskVariation:
		return &Disk{}, nil
	case BubbleVariation:
		return &Bubble{}, nil
	case HeartVariation:
		return &Heart{}, nil
	case SpiralVariation:
		return &Spiral{}, nil
	default:
		return nil, &UnknownVariationError{Name: name}
	}
}

type Transformation struct {
	Affine    AffineTransform
	Variation Transformer
	Color     [3]float64
	Weight    float64
}

type Transformations []Transformation

func (t Transformations) ChooseTransformation(rng *rand.Rand) Transformation {
	totalWeight := 0.0
	for _, t := range t {
		totalWeight += t.Weight
	}

	r := rng.Float64() * totalWeight
	sum := 0.0

	for _, t := range t {
		sum += t.Weight
		if r <= sum {
			return t
		}
	}

	return t[len(t)-1]
}

type AffineTransform struct {
	A, B, C, D, E, F float64
}

func (at *AffineTransform) Apply(p Point) Point {
	x := at.A*p.X + at.B*p.Y + at.C
	y := at.D*p.X + at.E*p.Y + at.F

	return Point{X: x, Y: y}
}

type Linear struct{}

func (l *Linear) Transform(p Point) Point {
	return p
}

type Sinusoidal struct{}

func (s *Sinusoidal) Transform(p Point) Point {
	return Point{
		X: math.Sin(p.X),
		Y: math.Sin(p.Y),
	}
}

type Spherical struct{}

func (s *Spherical) Transform(p Point) Point {
	r2 := p.X*p.X + p.Y*p.Y

	return Point{
		X: p.X / r2,
		Y: p.Y / r2,
	}
}

type Swirl struct{}

func (s *Swirl) Transform(p Point) Point {
	r2 := p.X*p.X + p.Y*p.Y
	sinr2 := math.Sin(r2)
	cosr2 := math.Cos(r2)

	return Point{
		X: p.X*sinr2 - p.Y*cosr2,
		Y: p.X*cosr2 + p.Y*sinr2,
	}
}

type Horseshoe struct{}

func (h *Horseshoe) Transform(p Point) Point {
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)

	return Point{
		X: (p.X - p.Y) * (p.X + p.Y) / r,
		Y: 2 * p.X * p.Y / r,
	}
}

type Fisheye struct{}

func (f *Fisheye) Transform(p Point) Point {
	r := 2 / (math.Sqrt(p.X*p.X+p.Y*p.Y) + 1)

	return Point{
		X: r * p.Y,
		Y: r * p.X,
	}
}

type Polar struct{}

func (p *Polar) Transform(point Point) Point {
	theta := math.Atan2(point.Y, point.X)
	r := math.Sqrt(point.X*point.X + point.Y*point.Y)

	return Point{
		X: theta / math.Pi,
		Y: r - 1.0,
	}
}

type Disk struct{}

func (d *Disk) Transform(p Point) Point {
	r := math.Pi * math.Sqrt(p.X*p.X+p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	return Point{
		X: (theta / math.Pi) * math.Sin(r),
		Y: (theta / math.Pi) * math.Cos(r),
	}
}

type Bubble struct{}

func (b *Bubble) Transform(p Point) Point {
	r2 := p.X*p.X + p.Y*p.Y
	scale := 4.0 / (r2 + 4.0)

	return Point{
		X: p.X * scale,
		Y: p.Y * scale,
	}
}

type Heart struct{}

func (h *Heart) Transform(p Point) Point {
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	return Point{
		X: r * math.Sin(theta*r),
		Y: -r * math.Cos(theta*r),
	}
}

type Spiral struct{}

func (s *Spiral) Transform(p Point) Point {
	r := math.Sqrt(p.X*p.X + p.Y*p.Y)
	theta := math.Atan2(p.Y, p.X)

	return Point{
		X: (1 / r) * (math.Cos(theta) + math.Sin(r)),
		Y: (1 / r) * (math.Sin(theta) - math.Cos(r)),
	}
}
