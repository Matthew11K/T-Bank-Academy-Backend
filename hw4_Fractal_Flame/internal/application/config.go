package application

import "github.com/central-university-dev/backend_academy_2024_project_4-go-Matthew11K/internal/domain"

type TransformationConfig struct {
	Type   string                 `json:"type"`
	Color  [3]float64             `json:"color"`
	Weight float64                `json:"weight"`
	Affine domain.AffineTransform `json:"affine"`
}

type Config struct {
	Width           int
	Height          int
	Iterations      int
	Points          int
	Threads         int
	Output          string
	Gamma           float64
	Symmetry        int
	Transformations domain.Transformations
}
